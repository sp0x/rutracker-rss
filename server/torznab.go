package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/sp0x/torrentd/indexer"
	"github.com/sp0x/torrentd/indexer/cache"
	"github.com/sp0x/torrentd/indexer/search"
	"github.com/sp0x/torrentd/torrent"
	"github.com/sp0x/torrentd/torznab"
)

func (s *Server) aggregatesStatus(c *gin.Context) {
	aggregate, err := s.indexerFacade.LoadedIndexes.Lookup(s.config, "all")
	if err != nil {
		_ = c.Error(err)
		return
	}

	statusObj := struct {
		ActiveIndexers []string `json:"latest"`
	}{}
	for _, ixr := range aggregate.(*indexer.Aggregate).Indexers {
		ixrInfo := ixr.Info()
		statusObj.ActiveIndexers = append(statusObj.ActiveIndexers, ixrInfo.GetTitle())
	}
	c.JSON(200, statusObj)
}

var searchCache, _ = cache.NewTTL(100, 1*time.Hour)

// Handle queries
func (s *Server) torznabHandler(c *gin.Context) {
	_ = c.Params
	indexerID := indexer.ResolveIndexID(s.indexerFacade.LoadedIndexes, c.Param("searchIndex"))
	t := c.Query("t")
	searchIndex, err := s.indexerFacade.LoadedIndexes.Lookup(s.config, indexerID)
	if err != nil {
		torznab.Error(c, err.Error(), torznab.ErrIncorrectParameter)
		return
	}
	if t == "caps" {
		searchIndex.Capabilities().ServeHTTP(c.Writer, c.Request)
		return
	}
	apiKey := c.Query("apikey")
	if !s.checkAPIKey(apiKey) {
		torznab.Error(c, "Invalid apikey parameter", torznab.ErrInsufficientPrivs)
		return
	}

	if t == "" {
		http.Redirect(c.Writer, c.Request, c.Request.URL.Path+"?t=caps", http.StatusTemporaryRedirect)
		return
	}

	switch t {
	case "search", "tvsearch", "tv-search", "movie", "movie-search", "moviesearch":
		query, err := search.ParseQuery(c.Request.URL.Query())
		if err != nil {
			torznab.Error(c, "Invalid query", torznab.ErrInsufficientPrivs)
			return
		}
		err = torrent.EnrichMovieAndShowData(query)
		if err != nil {
			torznab.Error(c, "Invalid query", torznab.ErrInsufficientPrivs)
			return
		}

		var feed *torznab.ResultFeed
		if cachedFeed, ok := searchCache.Get(query.UniqueKey()); ok {
			feed = cachedFeed.(*torznab.ResultFeed)
		} else {
			feed, err = s.torznabSearch(c.Request, query, s.indexerFacade)
			searchCache.Add(query.UniqueKey(), feed)
		}
		if err != nil {
			torznab.Error(c, err.Error(), torznab.ErrUnknownError)
			return
		}
		switch c.Query("format") {
		case "atom":
			atomOutput(c, feed)
		case "", "xml":
			xmlOutput(c, feed, searchIndex.GetEncoding())
		case "json":
			jsonOutput(c.Writer, feed, searchIndex.GetEncoding())
		}

	default:
		torznab.Error(c, "Unknown type parameter", torznab.ErrIncorrectParameter)
	}
}

func formatEncoding(nm string) string {
	nm = strings.ReplaceAll(nm, "ows12", "ows-12")
	nm = strings.Title(nm)
	return nm
}

func (s *Server) torznabSearch(r *http.Request, query *search.Query, indexer *indexer.Facade) (*torznab.ResultFeed, error) {
	srch := search.NewSearch(query)
	srch, err := indexer.Search(srch, query)
	if err != nil {
		return nil, err
	}
	nfo := indexer.Index.Info()

	feed := &torznab.ResultFeed{
		Info: torznab.Info{
			ID:          nfo.GetID(),
			Title:       nfo.GetTitle(),
			Description: "",
			Link:        nfo.GetLink(),
			Language:    nfo.GetLanguage(),
			Category:    "",
		},
		Items: srch.GetResults(),
	}
	feed.Info.Category = query.Type

	rewritten, err := s.rewriteLinks(r, srch.GetResults())
	if err != nil {
		return nil, err
	}

	feed.Items = rewritten

	return feed, err
}

// Rewrites the download links so that the download goes through us.
// This is required since only we can access the torrent ( the site might need authorization )
func (s *Server) rewriteLinks(r *http.Request, items []search.ResultItemBase) ([]search.ResultItemBase, error) {
	baseURL, err := s.baseURL(r, "/d")
	if err != nil {
		return nil, err
	}
	apiKey := s.sharedKey()
	// rewrite non-magnet links to use the server
	for idx, item := range items {
		scrapeItem := item.AsScrapeItem()
		if strings.HasPrefix(scrapeItem.Link, "magnet:") {
			continue
		}
		// itemTmp := item
		tokenValue, err := getTokenValue(scrapeItem, apiKey)
		if err != nil {
			return nil, err
		}
		filename := strings.ReplaceAll(item.UUID(), "/", "-")
		items[idx].AsScrapeItem().Link = fmt.Sprintf("%s/%s/%s.torrent", baseURL.String(), tokenValue, url.QueryEscape(filename))
	}

	return items, nil
}

func getTokenValue(item *search.ScrapeResultItem, apiKey []byte) (string, error) {
	sourceLink := item.SourceLink
	if sourceLink == "" {
		sourceLink = item.Link
	}
	indexerName := ""
	if item.Indexer != nil {
		indexerName = item.Indexer.Name
	} else {
		indexerName = item.Site
	}
	// Encode the site and source of the torrent as a JWT token
	t := &token{
		IndexName: indexerName,
		Link:      sourceLink,
	}

	te, err := t.Encode(apiKey)
	if err != nil {
		log.Debugf("Error encoding token: %v", err)
		return "", err
	}
	return te, nil
}

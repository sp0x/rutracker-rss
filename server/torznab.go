package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sp0x/rutracker-rss/indexer"
	"github.com/sp0x/rutracker-rss/indexer/search"
	"github.com/sp0x/rutracker-rss/torznab"
	"net/http"
	"net/url"
	"strings"
)

func (s *Server) torznabHandler(c *gin.Context) {
	_ = c.Params
	indexerID := c.Param("indexer")
	//Type of operation
	t := c.Query("t")
	indexer, err := s.lookupIndexer(indexerID)
	if err != nil {
		torznab.Error(c, err.Error(), torznab.ErrIncorrectParameter)
		return
	}
	switch t {
	case "caps":
		indexer.Capabilities().ServeHTTP(c.Writer, c.Request)
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
	case "caps":
		indexer.Capabilities().ServeHTTP(c.Writer, c.Request)

	case "search", "tvsearch", "tv-search", "movie", "movie-search", "moviesearch":
		feed, err := s.torznabSearch(c.Request, indexer, indexerID)
		if err != nil {
			torznab.Error(c, err.Error(), torznab.ErrUnknownError)
			return
		}
		switch c.Query("format") {
		case "atom":
			atomOutput(c, feed, indexer.GetEncoding())
		case "", "xml":
			xmlOutput(c, feed, indexer.GetEncoding())
		case "json":
			jsonOutput(c.Writer, feed, indexer.GetEncoding())
		}

	default:
		torznab.Error(c, "Unknown type parameter", torznab.ErrIncorrectParameter)
	}
}

func formatEncoding(nm string) string {
	nm = strings.Replace(nm, "ows12", "ows-12", -1)
	nm = strings.Title(nm)
	return nm
}

func (s *Server) createIndexer(key string) (torznab.Indexer, error) {
	def, err := indexer.DefaultDefinitionLoader.Load(key)
	if err != nil {
		log.WithError(err).Warnf("Failed to load definition for %q. %v", key, err)
		return nil, err
	}

	log.WithFields(log.Fields{"indexer": key}).Debugf("Loaded indexer")
	indexer, err := indexer.NewRunner(def, indexer.RunnerOpts{
		Config: s.Params.Config,
	}), nil
	if err != nil {
		return nil, err
	}

	return indexer, nil
}

func (s *Server) lookupIndexer(key string) (torznab.Indexer, error) {
	if key == "aggregate" || key == "all" {
		return s.createAggregate()
	}
	if _, ok := s.indexers[key]; !ok {
		indexer, err := s.createIndexer(key)
		if err != nil {
			return nil, err
		}
		s.indexers[key] = indexer
	}

	return s.indexers[key], nil
}

func (s *Server) torznabSearch(r *http.Request, indexer torznab.Indexer, siteKey string) (*torznab.ResultFeed, error) {
	query, err := torznab.ParseQuery(r.URL.Query())
	if err != nil {
		return nil, err
	}

	srch, err := indexer.Search(query)
	if err != nil {
		return nil, err
	}

	feed := &torznab.ResultFeed{
		Info:  indexer.Info(),
		Items: srch.Results,
	}
	feed.Info.Category = query.Type

	rewritten, err := s.rewriteLinks(r, srch.Results)
	if err != nil {
		return nil, err
	}

	feed.Items = rewritten

	return feed, err
}

func (s *Server) rewriteLinks(r *http.Request, items []search.ResultItem) ([]search.ResultItem, error) {
	baseURL, err := s.baseURL(r, "/download")
	if err != nil {
		return nil, err
	}

	k, err := s.sharedKey()
	if err != nil {
		return nil, err
	}

	// rewrite non-magnet links to use the server
	for idx, item := range items {
		if strings.HasPrefix(item.Link, "magnet:") {
			continue
		}
		//Encode the site and source of the torrent as a JWT token
		t := &token{
			Site: item.Site,
			Link: item.SourceLink,
		}

		te, err := t.Encode(k)
		if err != nil {
			log.Debugf("Error encoding token: %v", err)
			return nil, err
		}

		filename := strings.Replace(item.Title, "/", "-", -1)
		items[idx].Link = fmt.Sprintf("%s/%s/%s.torrent", baseURL.String(), te, url.QueryEscape(filename))
	}

	return items, nil
}

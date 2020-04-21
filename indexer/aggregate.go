package indexer

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/sp0x/rutracker-rss/indexer/search"
	"github.com/sp0x/rutracker-rss/torznab"
	"io"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type Aggregate []torznab.Indexer

func (ag Aggregate) GetEncoding() string {
	for _, indexer := range ag {
		return indexer.GetEncoding()
	}
	return "utf-8"
}

func (ag Aggregate) Search(query torznab.Query) (*search.Search, error) {
	g := errgroup.Group{}
	allResults := make([][]search.ResultItem, len(ag))
	maxLength := 0

	// fetch all results
	for idx, indexer := range ag {
		indexerID := indexer.Info().ID
		idx, indexer := idx, indexer
		g.Go(func() error {
			srchRes, err := indexer.Search(query)
			if err != nil {
				log.Warnf("Indexer %q failed: %s", indexerID, err)
				return nil
			}
			allResults[idx] = srchRes.Results
			if l := len(srchRes.Results); l > maxLength {
				maxLength = l
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		log.Warn(err)
		return nil, err
	}

	var outputSearch = &search.Search{}
	var results []search.ResultItem

	// interleave search results to preserve ordering
	for i := 0; i <= maxLength; i++ {
		for _, r := range allResults {
			if len(r) > i {
				results = append(results, r[i])
			}
		}
	}

	if query.Limit > 0 && len(results) > query.Limit {
		results = results[:query.Limit]
	}
	outputSearch.Results = results
	return outputSearch, nil
}

func (ag Aggregate) Info() torznab.Info {
	return torznab.Info{
		ID:       "aggregate",
		Title:    "Aggregated Indexer",
		Language: "en-US",
		Link:     "",
	}
}

func (ag Aggregate) Capabilities() torznab.Capabilities {
	return torznab.Capabilities{
		SearchModes: []torznab.SearchMode{
			{Key: "movie-search", Available: true, SupportedParams: []string{"q", "imdbid"}},
			{Key: "tv-search", Available: true, SupportedParams: []string{"q", "season", "ep"}},
			{Key: "search", Available: true, SupportedParams: []string{"q"}},
		},
	}
}

func (ag Aggregate) Download(u string) (io.ReadCloser, http.Header, error) {
	return nil, http.Header{}, errors.New("Not implemented")
}

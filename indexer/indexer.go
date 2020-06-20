package indexer

/**
This is part of https://github.com/cardigann/cardigann
*/

import (
	"github.com/sp0x/torrentd/indexer/search"
	"github.com/sp0x/torrentd/torznab"
	"io"
	"net/http"
)

type Info interface {
	GetId() string
	GetTitle() string
	GetLanguage() string
	GetLink() string
}

//go:generate mockgen -source indexer.go -destination=mocks/indexer.go -package=mocks
type Indexer interface {
	Info() Info
	Search(query torznab.Query, srch search.Instance) (search.Instance, error)
	Download(urlStr string) (io.ReadCloser, error)
	Capabilities() torznab.Capabilities
	GetEncoding() string
	ProcessRequest(req *http.Request) (*http.Response, error)
	Open(s *search.ExternalResultItem) (io.ReadCloser, error)
	//Check if the Indexer works ok.
	//This might be needed to validate the search result extraction.
	Check() error
	//The maximum number of pages we can search
	MaxSearchPages() uint
	SearchIsSinglePaged() bool
}

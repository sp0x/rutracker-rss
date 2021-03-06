package torznab

import (
	"encoding/xml"
	"net/http"
	"sort"
	"strings"

	"github.com/sp0x/torrentd/indexer/categories"
	"github.com/sp0x/torrentd/indexer/search"
)

type Capabilities struct {
	SearchModes []search.Capability
	Categories  categories.Categories
}

func (c Capabilities) HasCategory(cat categories.Category) bool {
	return c.Categories.ContainsCat(&cat)
}

func (c Capabilities) HasCategories(cats []categories.Category) bool {
	for _, theirCat := range cats {
		if !c.Categories.ContainsCat(&theirCat) {
			return false
		}
	}
	return true
}

func (c Capabilities) HasSearchMode(key string) (bool, []string) {
	for _, m := range c.SearchModes {
		if m.Key == key && m.Available {
			return true, m.SupportedParams
		}
	}
	return false, nil
}

func (c Capabilities) HasTVShows() bool {
	for _, cat := range c.Categories {
		if cat.ID >= 5000 && cat.ID < 6000 {
			return true
		}
	}
	return false
}

func (c Capabilities) HasMovies() bool {
	for _, cat := range c.Categories {
		if cat.ID >= 2000 && cat.ID < 3000 {
			return true
		}
	}
	return false
}

func (c Capabilities) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var cx struct {
		XMLName   struct{} `xml:"caps"`
		Searching struct {
			Values []interface{}
		} `xml:"searching"`
		Categories struct {
			Values []interface{}
		} `xml:"categories"`
	}

	for _, mode := range c.SearchModes {
		available := "no"
		if mode.Available {
			available = "yes"
		}
		cx.Searching.Values = append(cx.Searching.Values, struct {
			XMLName         xml.Name
			Available       string `xml:"available,attr"`
			SupportedParams string `xml:"supportedParams,attr"`
		}{
			xml.Name{Space: "", Local: mode.Key},
			available,
			strings.Join(mode.SupportedParams, ","),
		})
	}

	cats := c.Categories
	sort.Sort(cats)

	for _, cat := range cats {
		cx.Categories.Values = append(cx.Categories.Values, struct {
			XMLName struct{} `xml:"category"`
			ID      int      `xml:"id,attr"`
			Name    string   `xml:"name,attr"`
		}{
			ID:   cat.ID,
			Name: cat.Name,
		})
	}

	err := e.Encode(cx)
	return err
}

func (c Capabilities) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	x, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	_, _ = w.Write(x)
}

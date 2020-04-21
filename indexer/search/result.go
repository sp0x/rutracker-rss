package search

import (
	"encoding/xml"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

const rfc822 = "Mon, 02 Jan 2006 15:04:05 -0700"

type torznabAttrView struct {
	XMLName struct{} `xml:"torznab:attr"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type ExternalResultItem struct {
	gorm.Model
	ResultItem
	LocalCategoryID   string
	LocalCategoryName string
	LocalId           string
	IsMagnet          bool
	Announce          string
	Publisher         string
	Fingerprint       string
}

type ResultItem struct {
	Site          string
	Title         string
	OriginalTitle string
	Description   string
	GUID          string
	Comments      string
	Link          string

	SourceLink  string
	Category    int
	Size        uint64
	Files       int
	Grabs       int
	PublishDate int64

	Seeders              int
	Peers                int
	MinimumRatio         float64
	MinimumSeedTime      time.Duration
	DownloadVolumeFactor float64
	UploadVolumeFactor   float64
	Author               string
	AuthorId             string
	Indexer              *ResultIndexer
}

type ResultIndexer struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"` //make the name the value
}

func (ri *ResultItem) AddedOnStr() interface{} {
	tm := time.Unix(ri.PublishDate, 0)
	return tm.String()
}

func (ri ResultItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//The info view enclosure
	var enclosure = struct {
		URL    string `xml:"url,attr,omitempty"`
		Length uint64 `xml:"length,attr,omitempty"`
		Type   string `xml:"type,attr,omitempty"`
	}{
		URL:    ri.SourceLink,
		Length: ri.Size,
		Type:   "application/x-bittorrent",
	}
	var atomLink = struct {
		XMLName string `xml:"atom:link"`
		Href    string `xml:"href,attr"`
		Rel     string `xml:"rel,attr"`
		Type    string `xml:"type,attr"`
	}{
		Href: "http://jackett.vaskovasilev.eu/", Rel: "self", Type: "application/rss+xml",
	}
	var itemView = struct {
		XMLName  struct{} `xml:"item"`
		AtomLink interface{}
		// standard rss elements
		Title             string         `xml:"title,omitempty"`
		Indexer           *ResultIndexer `xml:"indexer,omitempty"`
		Description       string         `xml:"description,omitempty"`
		GUID              string         `xml:"guid,omitempty"`
		Comments          string         `xml:"comments,omitempty"`
		Link              string         `xml:"link,omitempty"`
		Category          string         `xml:"category,omitempty"`
		Files             int            `xml:"files,omitempty"`
		Grabs             int            `xml:"grabs,omitempty"`
		PublishDate       string         `xml:"pubDate,omitempty"`
		Enclosure         interface{}    `xml:"enclosure,omitempty"`
		Size              uint64         `xml:"size"`
		TorznabAttributes []torznabAttrView
	}{
		Title:       ri.Title,
		Description: ri.Description,
		Indexer:     ri.Indexer,
		GUID:        ri.GUID,
		Comments:    ri.Comments,
		Link:        ri.Link,
		Category:    strconv.Itoa(ri.Category),
		Files:       ri.Files,
		Grabs:       ri.Grabs,
		PublishDate: time.Unix(ri.PublishDate, 0).Format(rfc822),
		Enclosure:   enclosure,
		AtomLink:    atomLink,
		Size:        ri.Size,
	}
	attribs := itemView.TorznabAttributes
	attribs = append(attribs, torznabAttrView{Name: "category", Value: strconv.Itoa(ri.Category)})
	attribs = append(attribs, torznabAttrView{Name: "seeds", Value: strconv.Itoa(ri.Seeders)})
	attribs = append(attribs, torznabAttrView{Name: "peers", Value: strconv.Itoa(ri.Peers)})
	attribs = append(attribs, torznabAttrView{Name: "minimumratio", Value: fmt.Sprint(ri.MinimumRatio)})
	attribs = append(attribs, torznabAttrView{Name: "minimumseedtime", Value: fmt.Sprint(ri.MinimumSeedTime)})
	attribs = append(attribs, torznabAttrView{Name: "downloadvolumefactor", Value: fmt.Sprint(ri.DownloadVolumeFactor)})
	attribs = append(attribs, torznabAttrView{Name: "uploadvolumefactor", Value: fmt.Sprint(ri.UploadVolumeFactor)})

	itemView.TorznabAttributes = attribs
	e.Encode(itemView)
	return nil
}
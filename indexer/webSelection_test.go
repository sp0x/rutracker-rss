package indexer

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/onsi/gomega"
	"strings"
	"testing"
)

func Test_ShouldMatchTextForSimpleSelectors(t *testing.T) {
	g := gomega.NewWithT(t)
	selector := selectorBlock{Selector: "a"}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader("<div><a>Inner Text</a></div>"))
	if err != nil {
		t.Fail()
		return
	}
	if doc == nil {
		t.Fail()
		return
	}
	selection := &DomScrapeItem{doc.Contents()}
	result, err := selector.Match(selection)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(result).To(gomega.Equal("Inner Text"))
}

func Test_ShouldMatchTextForSelectorsWithMultipleMatches(t *testing.T) {
	g := gomega.NewWithT(t)
	selector := selectorBlock{Selector: "a", All: true}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader("<div><a>Inner Text</a><a>Other Text</a></div>"))
	if err != nil {
		t.Fail()
		return
	}
	if doc == nil {
		t.Fail()
		return
	}
	selection := doc.Contents()
	result, err := selector.Match(&DomScrapeItem{selection})
	g.Expect(err).To(gomega.BeNil())
	g.Expect(result.([]string)).ToNot(gomega.BeNil())
	resultArray := result.([]string)
	g.Expect(resultArray[0]).To(gomega.Equal("Inner Text"))
	g.Expect(resultArray[1]).To(gomega.Equal("Other Text"))
}

package indexer

import (
	"errors"
	"fmt"
	"text/template"

	"github.com/sp0x/torrentd/indexer/search"
	"github.com/sp0x/torrentd/indexer/utils"
)

type SearchTemplateData struct {
	Query      *search.Query
	Keywords   string
	Categories []string
	Context    *RunContext
	Functions  template.FuncMap
}

func (s *SearchTemplateData) ApplyTo(fieldName string, templateText string) (string, error) {
	fmap := s.Functions
	fmap["rng"] = func(start, end string) string {
		return s.RangeValue(fieldName)
	}
	return utils.ApplyTemplate(fieldName, templateText, s, s.Functions)
}

func (s *SearchTemplateData) HasQueryField(name string) bool {
	if s.Query == nil {
		return false
	}
	_, ok := s.Query.Fields[name]
	return ok
}

func (s *SearchTemplateData) ApplyField(fieldName string) (string, error) {
	if s.Query == nil {
		return "", errors.New("template has no query")
	}
	fieldValue := s.Query.Fields[fieldName]
	switch value := fieldValue.(type) {
	case search.RangeField:
		return s.RangeValue(fieldName), nil
	default:
		return fmt.Sprint(value), nil
	}
}

func (s *SearchTemplateData) RangeValue(name string) string {
	fieldStates := s.Context.Search.FieldState
	if _, ok := fieldStates[name]; !ok {
		return ""
	}
	fieldState := fieldStates[name]
	if !fieldState.HasNext() {
		return ""
	}
	nextValue := fieldState.Next()
	return nextValue
}

func newSearchTemplateData(query *search.Query, localCategories []string, context *RunContext) *SearchTemplateData {
	searchData := &SearchTemplateData{
		query,
		query.Keywords(),
		localCategories,
		context,
		utils.GetDefaultFunctionMap(),
	}
	return searchData
}

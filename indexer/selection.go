package indexer

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/yosssi/gohtml"
)

type filterBlock struct {
	Name string      `yaml:"name"`
	Args interface{} `yaml:"args"`
}

type selectorBlock struct {
	Selector  string            `yaml:"selector"`
	Pattern   string            `yaml:"pattern"`
	TextVal   string            `yaml:"text"`
	Attribute string            `yaml:"attribute,omitempty"`
	Remove    string            `yaml:"remove,omitempty"`
	Filters   []filterBlock     `yaml:"filters,omitempty"`
	Case      map[string]string `yaml:"case,omitempty"`
}

func (s *selectorBlock) Match(selection *goquery.Selection) bool {
	return !s.IsEmpty() && (selection.Find(s.Selector).Length() > 0 || s.TextVal != "")
}

func (s *selectorBlock) MatchText(from *goquery.Selection) (string, error) {
	if s.TextVal != "" {
		return s.TextVal, nil
	}
	if s.Selector != "" {
		result := from.Find(s.Selector)
		if result.Length() == 0 {
			return "", fmt.Errorf("Failed to match selector %q", s.Selector)
		}
		return s.Text(result)
	}
	if s.Pattern != "" {
		return s.Pattern, nil
	}
	return s.Text(from)
}

//Text extracts text from the selection
func (s *selectorBlock) Text(el *goquery.Selection) (string, error) {
	if s.TextVal != "" {
		return s.applyFilters(s.TextVal)
	}

	if s.Remove != "" {
		el.Find(s.Remove).Remove()
	}

	if s.Case != nil {
		filterLogger.WithFields(logrus.Fields{"case": s.Case}).
			Debugf("Applying case to selection")
		for pattern, value := range s.Case {
			if el.Is(pattern) || el.Has(pattern).Length() >= 1 {
				return s.applyFilters(value)
			}
		}
		return "", errors.New("None of the cases match")
	}

	html, _ := goquery.OuterHtml(el)
	filterLogger.
		WithFields(logrus.Fields{"html": gohtml.Format(html)}).
		Debugf("Extracting text from selection")

	output := strings.TrimSpace(el.Text())

	if s.Attribute != "" {
		val, exists := el.Attr(s.Attribute)
		if !exists {
			return "", fmt.Errorf("Requested attribute %q doesn't exist", s.Attribute)
		}
		output = val
	}

	return s.applyFilters(output)
}

//Filter the value through a list of filters
func (s *selectorBlock) applyFilters(val string) (string, error) {
	prevFilterFailed := false
	var prevFilter filterBlock
	for _, f := range s.Filters {
		shouldFilter := true
		switch f.Name {
		case "dateparseAlt":
			//This is ran only if there has been a failure before.
			shouldFilter = prevFilterFailed && prevFilter.Name == "dateparse"
			break
		default:
			shouldFilter = true
		}
		if !shouldFilter {
			continue
		}

		filterLogger.WithFields(logrus.Fields{"args": f.Args, "before": val}).
			Debugf("Applying filter %s", f.Name)
		var err error
		newVal, err := invokeFilter(f.Name, f.Args, val)
		if err != nil {
			if f.Name != "dateparse" {
				logrus.Warningf("Filter %s failed on value `%v`. %s\n", f.Name, val, err)
			}
			prevFilterFailed = true
			prevFilter = f
			continue
			//return "", err
		}
		val = newVal
		prevFilterFailed = false
	}

	return val, nil
}

func (s *selectorBlock) IsEmpty() bool {
	return s.Selector == "" && s.TextVal == ""
}

func (s *selectorBlock) String() string {
	switch {
	case s.Selector != "":
		return fmt.Sprintf("Selector(%s)", s.Selector)
	case s.TextVal != "":
		return fmt.Sprintf("Text(%s)", s.TextVal)
	default:
		return "Empty"
	}
}

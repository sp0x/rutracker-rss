package indexer

import (
	"github.com/sp0x/rutracker-rss/indexer/categories"
)

type categoryMap map[string]*categories.Category

//Categories gets the collection of categories that this map contains
func (mapping categoryMap) Categories() categories.Categories {
	cats := categories.Categories{}
	added := map[int]bool{}

	for _, c := range mapping {
		if _, exists := added[c.ID]; exists {
			continue
		}
		cats[c.ID] = c
		added[c.ID] = true
	}

	return cats
}

func (mapping categoryMap) Resolve(cat *categories.Category) []string {
	var matched bool
	var results = []string{}

	// check for exact matches
	for localID, mappedCat := range mapping {
		if mappedCat.ID == cat.ID {
			results = append(results, localID)
			matched = true
		}
	}

	// check for matches on the parent categories of the mapped categories
	// e.g asked for Movies, but only had a more specific mapping for Movies/Blu-ray
	if !matched {
		for localID, mappedCat := range mapping {
			if categories.ParentCategory(mappedCat).ID == cat.ID {
				results = append(results, localID)
				matched = true
			}
		}
	}

	// finally check for matches on the parent category of the requested cat
	// e.g. asked for Movies/Blu-ray but no mapping, so try Movies instead
	if !matched {
		parent := categories.ParentCategory(cat)
		for localID, mappedCat := range mapping {
			if mappedCat.ID == parent.ID {
				results = append(results, localID)
				matched = true
			}
		}
	}

	return results
}

func (mapping categoryMap) ResolveAll(cats ...*categories.Category) []string {
	var results []string
	for _, cat := range cats {
		results = append(results, mapping.Resolve(cat)...)
	}

	return results
}

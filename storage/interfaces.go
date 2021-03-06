package storage

import (
	"github.com/sp0x/torrentd/indexer/search"
	"github.com/sp0x/torrentd/storage/indexing"
	"github.com/sp0x/torrentd/storage/stats"
)

type ItemStorage interface {
	Size() int64
	Find(query indexing.Query, output *search.ScrapeResultItem) error
	Add(item search.Record) error
	AddUniqueIndex(key *indexing.Key)
	NewWithKey(pk *indexing.Key) ItemStorage
	Close()
	SetKey(index *indexing.Key) error
	GetLatest(count int) []search.ResultItemBase
	ForEach(callback func(record search.Record))
	GetStats(showDebugInfo bool) *stats.Stats
	Truncate() error
}

type ItemStorageBacking interface {
	// Find tries to find a single record matching the query.
	Find(query indexing.Query, result interface{}) error
	HasIndex(meta *indexing.IndexMetadata) bool
	GetIndexes() map[string]indexing.IndexMetadata
	Update(query indexing.Query, item interface{}) error
	// CreateWithID creates a new record using a custom key
	CreateWithID(parts *indexing.Key, item search.Record, uniqueIndexKeys *indexing.Key) error
	// Create a new record with the default key (UUIDValue)
	Create(item search.Record, additionalPK *indexing.Key) error
	// Size is the size of the storage, as in records count
	Size() int64
	// GetLatest returns the latest `count` of records.
	GetLatest(count int) []search.ResultItemBase
	Close()
	ForEach(callback func(record search.Record))
	GetStats(showDebugInfo bool) *stats.Stats
	Truncate() error
}

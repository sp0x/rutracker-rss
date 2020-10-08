package bolt

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/sp0x/torrentd/storage/indexing"
)

type Metadata struct {
	Indexes map[string]indexing.IndexMetadata `json:"indexes"`
}

func (b *BoltStorage) setupMetadata() error {
	return b.Database.Update(func(tx *bolt.Tx) error {
		rootBucket, err := b.assertBucket(tx, resultsBucket)
		if err != nil {
			return err
		}
		bucket, err := b.assertNamespaceBucket(tx, resultsBucket)
		if err != nil {
			return err
		}
		b.loadGlobalMetadata(rootBucket)
		b.loadMetadata(bucket)
		return nil
	})
}

func (b *BoltStorage) saveMetadata(bucket *bolt.Bucket) {
	metadataBytes, _ := json.Marshal(b.metadata)
	_ = bucket.Put([]byte(metaBucket), metadataBytes)
}

func (b *BoltStorage) loadMetadata(bucket *bolt.Bucket) {
	metadataBytes := bucket.Get([]byte(metaBucket))
	metadata := &Metadata{}
	if metadataBytes != nil {
		err := json.Unmarshal(metadataBytes, &metadata)
		if err != nil {
			panic(fmt.Sprintf("couldn't load metadata: %v", err))
		}
	}
	b.metadata = metadata
}

func (b *BoltStorage) GetIndexes() map[string]indexing.IndexMetadata {
	if b.metadata == nil {
		return nil
	}
	return b.metadata.Indexes
}

func (b *BoltStorage) HasIndex(meta *indexing.IndexMetadata) bool {
	_, found := b.metadata.Indexes[meta.Name]
	return found
}

func (b *BoltStorage) HasIndexWithName(name string) bool {
	_, found := b.metadata.Indexes[name]
	return found
}

func (m *Metadata) AddIndex(name string, location string, isUnique bool) bool {
	if m.Indexes == nil {
		m.Indexes = make(map[string]indexing.IndexMetadata)
	}
	existingIx, exists := m.Indexes[name]
	if exists && existingIx.Location == location && existingIx.Unique == isUnique {
		return false
	}
	m.Indexes[name] = indexing.IndexMetadata{
		Name:     name,
		Unique:   isUnique,
		Location: location,
	}
	return true
}

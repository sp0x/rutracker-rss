package storage

import (
	"github.com/sp0x/torrentd/indexer/search"
	"reflect"
)

type KeyedStorage struct {
	//backing *DBStorage
	backing ItemStorageBacking
	key     Key
}

type Key []string

//TODO: Add limit, reverse order
type Query map[string]interface{}

//NewKey creates a new key using an array of fields.
func NewKey(fieldNames ...string) Key {
	var key Key
	for _, item := range fieldNames {
		key = append(key, item)
	}
	return key
}

//NewKeyedStorage creates a new keyed storage with the default storage backing.
func NewKeyedStorage(keyFields Key) *KeyedStorage {
	return &KeyedStorage{
		key:     keyFields,
		backing: DefaultStorageBacking(),
	}
}

//NewKeyedStorageWithBacking creates a new keyed storage with a custom storage backing.
func NewKeyedStorageWithBacking(key Key, storage ItemStorageBacking) *KeyedStorage {
	return &KeyedStorage{
		key:     key,
		backing: storage,
	}
}

//Add handles the discovery of the result, adding additional information like staleness state.
func (s *KeyedStorage) Add(item *search.ExternalResultItem) (bool, bool) {
	var existingResult *search.ExternalResultItem
	existingKey := s.GetKeyQueryFromItem(item)
	if existingKey != nil {
		existingResult = s.backing.Find(existingKey)
	}
	//if item.LocalId != "" {
	//	existingResult = defaultStorage.FindById(item.LocalId)
	//} else {
	//	existingResult = defaultStorage.FindByNameAndIndex(item.Title, item.Site)
	//}
	isNew := existingResult == nil || existingResult.PublishDate != item.PublishDate
	isUpdate := existingResult != nil && (existingResult.PublishDate != item.PublishDate)
	if isNew {
		if isUpdate && existingResult != nil {
			item.Fingerprint = existingResult.Fingerprint
			s.backing.Update(existingKey, item)
			//defaultStorage.UpdateResult(existingResult.ID, item)
		} else {
			item.Fingerprint = search.GetResultFingerprint(item)
			//defaultStorage.Create(item)
			s.backing.Create(item)
		}
	}
	//We set the result's state so it's known later on whenever it's used.
	item.SetState(isNew, isUpdate)
	return isNew, isUpdate
}

//GetKeyQueryFromItem gets the query that matches an item with the given key.
func (s *KeyedStorage) GetKeyQueryFromItem(item *search.ExternalResultItem) Query {
	output := Query{}
	val := reflect.ValueOf(item).Elem()
	for _, kfield := range s.key {
		fld := val.FieldByName(kfield)
		if !fld.IsValid() {
			output[kfield] = item.GetField(kfield)
		} else {
			output[kfield] = fld.Interface()
		}
	}
	return output
}

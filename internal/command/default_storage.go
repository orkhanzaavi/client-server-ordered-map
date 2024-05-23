package command

import "sync"

// defaultStorage is a default implementation of the Storage interface
// it implements ordered hashmap to keep items in memory in the strict order they were added
type defaultStorage struct {
	// rootItem is the first item in the list
	// it is used to get all items in the order they were added
	rootItem *StoredItem
	// lastItem is the last item in the list
	// it is used to add new items to the end of the list with O(1) complexity
	lastItem *StoredItem
	values   map[string]*StoredItem

	mu sync.RWMutex
}

type StoredItem struct {
	Key   string
	Value string

	// make a two-way linked list to keep the order of items and delete them in O(1) complexity
	next *StoredItem
	prev *StoredItem
}

func newDefaultStorage() *defaultStorage {
	return &defaultStorage{
		values: make(map[string]*StoredItem),
	}
}

func (o *defaultStorage) Set(key, value string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	storedItem := o.Get(key)
	if storedItem == nil {
		storedItem = &StoredItem{
			Key:   key,
			Value: value,
		}
		o.add(storedItem)
		return
	}
	storedItem.Value = value
}

func (o *defaultStorage) add(storedItem *StoredItem) {
	if o.rootItem == nil {
		o.rootItem = storedItem
		o.lastItem = storedItem
	} else {
		storedItem.prev = o.lastItem
		o.lastItem.next = storedItem
		o.lastItem = storedItem
	}

	o.values[storedItem.Key] = storedItem
}

func (o *defaultStorage) Get(key string) *StoredItem {
	value, ok := o.values[key]
	if ok {
		return value
	}
	return nil
}

func (o *defaultStorage) GetAll() []*StoredItem {
	result := make([]*StoredItem, 0)
	if o.rootItem == nil {
		return result
	}
	item := o.rootItem
	for {
		result = append(result, item)
		if item.next == nil {
			break
		}
		item = item.next
	}
	return result
}

func (o *defaultStorage) Delete(key string) {
	item := o.Get(key)

	if item == nil {
		return
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	if item == o.rootItem {
		o.rootItem = item.next
	}
	if item == o.lastItem {
		o.lastItem = item.prev
	}
	if item.prev != nil {
		item.prev.next = item.next
	}
	if item.next != nil {
		item.next.prev = item.prev
	}

	delete(o.values, key)
}

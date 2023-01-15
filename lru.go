package cache_me

import (
	"container/list"
	"fmt"
)

type LRU struct {
	bc    *Base
	order *list.List
}

func NewLRU(sz int) *LRU {
	return &LRU{
		order: list.New(),
		bc:    newBase(sz),
	}
}

// Used to write a new entry or update exisiting one
func (cash *LRU) Put(key, value interface{}) error {
	item := &CacheItem{
		key:   key,
		value: value,
	}

	cash.bc.mu.Lock()
	defer cash.bc.mu.Unlock()

	ref, ok := cash.bc.store[key]

	if ok {
		// already exists

		cash.order.MoveToFront(ref)
		ref.Value.(*CacheItem).value = value

	} else {
		//size exceeded

		if cash.order.Len() >= cash.bc.len {
			last := cash.order.Back()
			// remove last element from list
			if last != nil {
				cash.RemoveElement(last)
			}
		}

		// point to order location in store
		cash.bc.store[key] = cash.order.PushFront(item)
	}

	return nil
}

// CANNOT USE READ LOCK because we are moving to front also here
func (cash *LRU) Get(key interface{}) (interface{}, error) {
	cash.bc.mu.Lock()
	ref, ok := cash.bc.store[key]

	if ok { // cache hit
		obj := ref.Value.(*CacheItem)
		// move to front because recently accessed
		cash.order.MoveToFront(ref)
		cash.bc.mu.Unlock()

		cash.bc.stat.incrementHit()
		return obj.value, nil
	}

	// cache miss
	cash.bc.mu.Unlock()
	cash.bc.stat.incrementMiss()
	return nil, ErrKeyNotFound
}

// Deletes element from map and list
func (cash *LRU) RemoveElement(ele *list.Element) {
	delete(cash.bc.store, ele.Value.(*CacheItem).key)
	cash.order.Remove(ele)
}

// Return current length of cache
func (cash *LRU) GetLength() int {
	cash.bc.mu.RLock()
	defer cash.bc.mu.RUnlock()
	return cash.order.Len()
}

// Fetch all stats
func (cash *LRU) GetStats() {
	fmt.Println("Hit Count : ", cash.bc.stat.HitCountFetch())
	fmt.Println("Miss Count : ", cash.bc.stat.MissCountFetch())
	fmt.Println("Lookup Count : ", cash.bc.stat.LookupCountFetch())
}

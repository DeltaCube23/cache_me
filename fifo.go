package cache_me

import (
	"container/list"
	"fmt"
)

type FIFO struct {
	bc    *Base
	order *list.List
}

func NewFIFO(sz int) *FIFO {
	return &FIFO{
		order: list.New(),
		bc:    newBase(sz),
	}
}

// Used to write a new entry or update exisiting one
func (cash *FIFO) Put(key, value interface{}) error {
	item := &CacheItem{
		key:   key,
		value: value,
	}

	cash.bc.mu.Lock()
	defer cash.bc.mu.Unlock()

	ref, ok := cash.bc.store[key]

	if ok {
		// already exists just update

		ref.Value.(*CacheItem).value = value
	} else {
		//size exceeded

		if cash.order.Len() >= cash.bc.len {
			last := cash.order.Front()
			// remove first element from list
			if last != nil {
				cash.RemoveElement(last)
			}
		}
		// point to order location in store
		cash.bc.store[key] = cash.order.PushBack(item)
	}

	return nil
}

// CAN USE READ LOCK
func (cash *FIFO) Get(key interface{}) (interface{}, error) {
	cash.bc.mu.RLock()
	ref, ok := cash.bc.store[key]

	if ok { // cache hit
		obj := ref.Value.(*CacheItem)
		cash.bc.mu.RUnlock()

		cash.bc.stat.incrementHit()
		return obj.value, nil
	}

	// cache miss
	cash.bc.mu.RUnlock()
	cash.bc.stat.incrementMiss()
	return nil, ErrKeyNotFound
}

// Deletes element from map and list
func (cash *FIFO) RemoveElement(ele *list.Element) {
	delete(cash.bc.store, ele.Value.(*CacheItem).key)
	cash.order.Remove(ele)
}

// Return current length of cache
func (cash *FIFO) GetLength() int {
	cash.bc.mu.RLock()
	defer cash.bc.mu.RUnlock()
	return cash.order.Len()
}

// Fetch all stats
func (cash *FIFO) GetStats() {
	fmt.Println("Hit Count : ", cash.bc.stat.HitCountFetch())
	fmt.Println("Miss Count : ", cash.bc.stat.MissCountFetch())
	fmt.Println("Lookup Count : ", cash.bc.stat.LookupCountFetch())
}

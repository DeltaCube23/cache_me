package cache_me

import (
	"container/list"
	"errors"
	"sync"
)

// Base Structure for Cache
type Cache interface {
	Put(key, value interface{}) error
	Get(key interface{}) (interface{}, error)
	RemoveElement(ele *list.Element)
	GetLength() int
	GetStats()
}

type Base struct {
	store map[interface{}]*list.Element
	len   int
	mu    sync.RWMutex
	stat  *Stats
}

func newBase(sz int) *Base {
	return &Base{
		store: make(map[interface{}]*list.Element),
		len:   sz,
		stat:  newStats(),
	}
}

type CacheItem struct {
	key   interface{}
	value interface{}
}

var ErrKeyNotFound = errors.New("key not found")

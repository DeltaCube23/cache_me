package cache_me

import (
	"container/list"
	"errors"
	"sync"
)

// Base Structure for Cache
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

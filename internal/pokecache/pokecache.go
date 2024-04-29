package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	duration time.Duration
	mux      *sync.RWMutex
}

func NewCache(cacheDuration time.Duration) Cache {
	newCache := Cache{
		duration: cacheDuration,
		entries:  make(map[string]cacheEntry),
		mux:      &sync.RWMutex{},
	}
	go newCache.reapLoop()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.duration)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mux.Lock()
	defer c.mux.Unlock()
	maxCacheTime := time.Now().UTC().Add(-c.duration)

	for key, entry := range c.entries {
		if entry.createdAt.Before(maxCacheTime) {
			delete(c.entries, key)
		}
	}
}

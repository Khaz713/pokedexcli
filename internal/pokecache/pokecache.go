package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Time{},
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.entries[key]
	if ok == false {
		return nil, false
	}
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cutoff := now.Add(-last)
	for k, v := range c.entries {
		if v.createdAt.Before(cutoff) {
			delete(c.entries, k)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		mu:      sync.Mutex{},
		entries: make(map[string]CacheEntry),
	}
	go newCache.reapLoop(interval)
	return newCache
}

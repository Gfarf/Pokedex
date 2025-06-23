package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mux  *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data[key] = cacheEntry{createdAt: time.Now().UTC(), value: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.data[key]
	return val.value, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.data, k)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		data: make(map[string]cacheEntry),
		mux:  &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

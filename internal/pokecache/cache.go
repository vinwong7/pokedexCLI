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
	cacheMap map[string]cacheEntry
	mu       *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}
	go c.ReapLoop(interval)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}
	return elem.val, true

}

func (c *Cache) ReapLoop(intveral time.Duration) {

	ticker := time.NewTicker(intveral)

	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.cacheMap {
			if time.Since(v.createdAt) >= intveral {
				delete(c.cacheMap, k)
			}
		}
		c.mu.Unlock()
	}
}

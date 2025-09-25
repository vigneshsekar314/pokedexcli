package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cacheMap: map[string]cacheEntry{},
	}
	tick := time.NewTicker(interval)
	go func() {
		for {
			<-tick.C
			go newCache.reapLoop(interval)
		}
	}()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	cache, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}
	c.mu.RUnlock()
	return cache.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	for key, value := range c.cacheMap {
		if (value.createdAt.Add(interval)).After(time.Now()) {
			delete(c.cacheMap, key)
		}
	}
	c.mu.Unlock()
}

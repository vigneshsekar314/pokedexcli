package pokecache

import (
	// "fmt"
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
		mu:       &sync.RWMutex{},
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
	// fmt.Printf("Adding cache key: %s and write locking map in Add()\n", key)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	// fmt.Printf("write unlocking map in Add()\n")
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// fmt.Printf("Getting cached key:%s. Read locking map\n", key)
	c.mu.RLock()
	defer c.mu.RUnlock()
	// fmt.Printf("Read locked map\n")
	cache, ok := c.cacheMap[key]
	if !ok {
		// fmt.Printf("Cache is empty, returning nil bytes and false. Read unlocked by defer statement\n")
		return nil, false
	}
	// fmt.Printf("Read Unlocked map, returning value from cache\n")
	return cache.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	// fmt.Printf("Write locking for reapLoop\n")
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.cacheMap {
		timenow := time.Now()
		if (value.createdAt.Add(interval)).Before(timenow) {
			// fmt.Printf("deleting as interval expired. interval: %v, created at: %v and time now: %v\n", interval, value.createdAt.Second(), timenow.Second())
			// fmt.Printf("deleting key: %s in reapLoop\n", key)
			delete(c.cacheMap, key)
		}
	}
	// fmt.Printf("Write Unlocking for reapLoop\n")
}

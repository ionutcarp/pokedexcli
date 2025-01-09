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
	mutex         sync.Mutex
	cacheLifetime time.Duration
	cacheEntries  map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntries:  make(map[string]cacheEntry),
		cacheLifetime: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, found := c.cacheEntries[key]
	if !found || time.Since(entry.createdAt) > c.cacheLifetime {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.cacheLifetime)
	for {
		<-ticker.C
		c.mutex.Lock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > c.cacheLifetime {
				delete(c.cacheEntries, key)
			}
		}
		c.mutex.Unlock()
	}
}

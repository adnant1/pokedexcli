package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	data []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache {
		cache: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, data []byte) {
	c.cache[key] = cacheEntry{
		data: data,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, found := c.cache[key]
	return entry.data, found
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	t := time.Now().UTC().Add(-interval)

	for key, entry := range c.cache {
		if entry.createdAt.Before(t) {
			delete(c.cache, key)
		}
	}
}
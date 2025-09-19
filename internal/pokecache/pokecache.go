package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	data []byte
	createdAt time.Time
}

func NewCache() Cache {
	return Cache {
		cache: make(map[string]cacheEntry),
	}
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
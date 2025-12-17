package pokecache


import (
	"sync"
	"time"
)


type cacheEntry struct{
	createdAt	time.Time
	val			[]byte
}

type Cache struct{
	sync.Mutex
	cache	map[string]cacheEntry
}


// NewCache creates a new chace with a configurable interval
func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache:  make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return &c
}


// reapLoop deletes entries older than the interval
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for t := range ticker.C{
		c.reap(t, interval)
	}
}


func (c *Cache) Add(key string, value []byte){
	c.Lock()
	defer c.Unlock()
	c.cache[key] = cacheEntry{
		val: value,
		createdAt: time.Now(),
	}
}


func (c *Cache) Get(key string) ([]byte, bool){
	c.Lock()
	defer c.Unlock()
	value, exists := c.cache[key]
	if !exists{
		return nil, false
	}
	return value.val, true
}


func (c *Cache) reap(now time.Time, last time.Duration){
	c.Lock()
	defer c.Unlock()
	for k, v := range c.cache{
		age := now.Sub(v.createdAt)
		if age > last {
			delete(c.cache, k)
		}
	}
}

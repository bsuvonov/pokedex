package pokecache

import (
	"fmt"
	"sync"
	"time"
)


type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}


type Cache struct {
	cache	 map[string]cacheEntry
	interval time.Duration
	mu		 *sync.Mutex
}


func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry {
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Lock()
	c.cache[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, found := c.cache[key]; !found {
		return nil, false
	} else {
		return val.val, true
	}
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.interval)
	for _ = range ticker.C {
		c.mu.Lock()
		for key, value := range c.cache {
			if time.Since(value.createdAt) >= c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) ReadCache() {
	ticker := time.NewTicker(1000*time.Millisecond)
	for _ = range ticker.C {
		c.mu.Lock()
		fmt.Print("\nin cache:")
		for key, _ := range c.cache {
			fmt.Print("\n" + key)
		}
		fmt.Println()
		c.mu.Unlock()
	}
}


func NewCache(interval time.Duration) Cache {
	cache := Cache{cache: make(map[string]cacheEntry), interval: interval, mu: &sync.Mutex{}}
	go cache.ReapLoop()
	return cache
}


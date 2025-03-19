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
	data map[string]cacheEntry
	mu   *sync.Mutex
}

// - New cache
func NewCache(interval time.Duration) Cache {
	c := Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

// - Add
func (c *Cache) Add(key string, newVal []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		val:       newVal,
		createdAt: time.Now().UTC(),
	}

}

// - Get
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	return val.val, ok

}

// - reapLoop -- создаёт тикер, который срабатывает через заданный интервал
// при срабатывании запускает функцию reap
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

// - reap -- защищает данные мьютексом, в цикле проходится по парам (ключ-значение)
// мапы c.data и если данные находятся в кэш дольше интервала, то эти данные удаляются
// из кэш
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.data, k)
		}
	}
}

package lcache

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type Cache struct {
	sync.Mutex
	cache      map[string]string
	accessTime map[int64]string
	size       int
}

type CacheResult struct {
	Value interface{}
	Error error
}

var KeyNotFoundError = errors.New("Key not found")

func NewCache(size ...int) *Cache { //make size optional
	data := make(map[string]string)
	timestamp := make(map[int64]string)
	if len(size) == 0 {
		return &Cache{
			Mutex:      sync.Mutex{},
			cache:      data,
			accessTime: timestamp,
			size:       -1,
		}
	}

	return &Cache{
		Mutex:      sync.Mutex{},
		cache:      data,
		accessTime: timestamp,
		size:       size[0],
	}

}
func (c *Cache) Read(key string) *CacheResult {
	c.Lock()
	defer c.Unlock()
	var r *CacheResult

	var value, keyExists = c.lookup(key)

	if !keyExists {
		return &CacheResult{
			Value: nil,
			Error: KeyNotFoundError,
		}
	}

	r = &CacheResult{
		Value: value,
		Error: nil,
	}
	//update last access time
	c.accessTime[time.Now().Unix()] = key

	return r
}

func (c *Cache) Write(key, data string) {
	c.Lock()
	defer c.Unlock()

	timestamp := time.Now().Unix()

	if c.size == -1 || len(c.cache) < c.size { //we haven't reached Cache size
		c.cache[key] = data
		c.accessTime[timestamp] = key
		return
	}

	if len(c.cache) == c.size {
		lru := c.getLeastAccessedKey(c.accessTime)
		delete(c.cache, lru)
	}

	//set or update
	c.cache[key] = data
	c.accessTime[timestamp] = key

	return

}
func (c *Cache) lookup(key string) (v string, ok bool) {
	v, ok = c.cache[key]
	return
}

//lRU check
func (c *Cache) getLeastAccessedKey(timestampMap map[int64]string) string {
	//return key corresponding to the earliest updated timestamp
	var timeEntries []int64
	for k, _ := range timestampMap {
		timeEntries = append(timeEntries, k)
	}

	sort.Slice(timeEntries, func(i, j int) bool {
		return timeEntries[i] < timeEntries[j]
	})

	return timestampMap[timeEntries[0]]
}

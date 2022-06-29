package lcache

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

//TODO

func TestReadCaches(t *testing.T) {

	type TC struct {
		name      string
		key       string
		data      string
		result    *CacheResult
		skipWrite bool
	}

	testcases := []TC{
		{
			name:      "ReadNonExisting",
			skipWrite: true,
			result: &CacheResult{
				Value: nil,
				Error: KeyNotFoundError,
			},
		},
		{
			name:      "ReadOk",
			key:       "key",
			data:      `{"a":1,"b":2}`,
			skipWrite: false,
			result: &CacheResult{
				Value: "{\"a\":1,\"b\":2}",
				Error: nil,
			},
		},
	}

	for _, tc := range testcases {
		testFunc := func(t *testing.T) {

			cache := NewCache()

			if !tc.skipWrite {
				cache.Write(tc.key, tc.data)
			}

			result := cache.Read(tc.key)

			assert.Equal(t, tc.result, result)

		}
		t.Run(tc.name, testFunc)

	}

}

func TestCachesSizeLimiting(t *testing.T) {

	type TC struct {
		size   int
		name   string
		key    string
		data   string
		result *CacheResult
	}

	testcases := []TC{

		{

			name: "TestCacheSizeLimit",
			size: 4,
			key:  "key",
			data: `{"a":1,"b":2}`,
			result: &CacheResult{
				Value: "{\"a\":1,\"b\":2}",
				Error: nil,
			},
		},
	}

	for _, tc := range testcases {
		testFunc := func(t *testing.T) {

			excess:=5
			cache := NewCache(tc.size)

			for i := 0; i < tc.size+excess; i++ {
				cache.Write(tc.key+"_"+strconv.Itoa(i), tc.data)
				time.Sleep(50 * time.Millisecond)
			}

			assert.Equal(t, tc.size, cache.Size())

		}
		t.Run(tc.name, testFunc)

	}

}

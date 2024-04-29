package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestEmptyCache(t *testing.T) {
	cache := NewCache(time.Minute)
	if cache.entries == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	cacheDuration := 5 * time.Minute
	cases := []struct {
		name string
		key  string
		val  []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("data2"),
		},
	}

	for _, cs := range cases {
		t.Run(fmt.Sprintf("Test case for key %v", cs.key), func(t *testing.T) {
			cache := NewCache(cacheDuration)
			cache.Add(cs.key, cs.val)
		})
	}
}

func TestReap(t *testing.T) {
	cacheDuration := 5 * time.Nanosecond
	key := "key"
	value := []byte("value")

	cache := NewCache(cacheDuration)
	cache.Add(key, value)

	time.Sleep(cacheDuration + time.Nanosecond)
	_, ok := cache.Get(key)
	if ok {
		t.Error("cache entry should have been reaped")
	}
}

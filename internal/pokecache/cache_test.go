package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "key1",
			val: []byte("val1"),
		},
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
	}

	for _, c := range cases {
		cache.Add(c.key, c.val)
		actual, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("%s not found", c.key)
			continue
		}
		if string(actual) != string(c.val) {
			t.Errorf("%s doesn't match %s", string(actual), string(c.val))
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key := "key1"
	cache.Add(key, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("%s should have been reaped", key)
	}
}
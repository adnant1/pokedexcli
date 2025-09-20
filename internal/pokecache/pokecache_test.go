package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("expected cache map to be initialized, got nil")
	}
}

func TestAddAndGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		key string
		data []byte
	}{
		{"key1", []byte("data1")},
		{"key2", []byte("data2")},
		{"key3", []byte("data3")},
	}

	for _, c := range cases {
		cache.Add(c.key, c.data)
		actual, ok := cache.Get(c.key)

		if !ok {
			t.Errorf("expected to find key %s in cache, but it was not found", c.key)
		}

		if string(actual) != string(c.data) {
			t.Errorf("for key %s, expected data %s, got %s", c.key, c.data, actual)
		}
	}
}

func TestReap (t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("key1", []byte("data1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get("key1")
	if ok {
		t.Error("expected key1 to be reaped, but it was found")
	}
}
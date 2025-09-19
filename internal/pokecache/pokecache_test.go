package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("expected cache map to be initialized, got nil")
	}
}

func TestAddAndGetCache(t *testing.T) {
	cache := NewCache()

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
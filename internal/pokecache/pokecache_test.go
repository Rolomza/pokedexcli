package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "key1",
			val: []byte("testdata"),
		},
		{
			key: "key2",
			val: []byte("moretestdata"),
		},
		{
			key: "",
			val: []byte("testdata"),
		},
	}

	for i, cas := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(cas.key, cas.val)
			val, ok := cache.Get(cas.key)
			if !ok {
				t.Errorf("expected to find key: %v", cas.key)
				return
			}

			if string(val) != string(cas.val) {
				t.Errorf("%s doesn't match %s", 
					string(val),
					cas.val)
				return
			}
		})
		
	}
	
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCache(baseTime)

	testKey := "https://example.com"

	cache.Add(testKey, []byte("testdata"))

	_, ok := cache.Get(testKey)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(testKey)
	if ok {
		t.Errorf("%s should have been reaped", testKey)
		return
	}
}
package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestPokeCacheAddGet(t *testing.T) {
	const interval = 10 * time.Millisecond
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/locations",
			val: []byte("some location areas"),
		},
		{
			key: "https://pokeapi.co/api/v2/locations/2",
			val: []byte("some more areas"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v for test key %v", i, c.key), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("could not find key %v", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("got %v, want %v", string(val), string(c.val))
				return
			}
		})
	}
}

func TestPokeCacheReapLoop(t *testing.T) {
	const interval = 10 * time.Millisecond
	cache := NewCache(interval)

	cache.Add("https://pokeapi.co/api/v2/locations", []byte("some location areas"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/locations")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(interval + 5*time.Millisecond)

	val, ok := cache.Get("https://pokeapi.co/api/v2/locations")
	if ok {
		t.Errorf("expected to not find key, found %v", string(val))
		return
	}
}

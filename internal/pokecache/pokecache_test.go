package pokecache

import (
	"testing"
	"time"
)

func TestReapLoop(t *testing.T) {

	cache := NewCache(5000*time.Millisecond)
	time.Sleep(2000*time.Millisecond)
	cache.Add("https://example.com", []byte("testdata"))

	time.Sleep(4000*time.Millisecond)
	cache.Add("https://example2.com", []byte("testdata2"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(5000*time.Millisecond)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
	}
	_, ok = cache.Get("https://example2.com")
	if !ok {
		t.Errorf("expected to find key")
	}

	time.Sleep(5000*time.Millisecond)
	_, ok = cache.Get("https://example2.com")
	if ok {
		t.Errorf("expected to not find key")
	}
}
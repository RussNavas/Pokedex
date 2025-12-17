package pokeapi

import (
	"testing"
	"time"
)

func TestListLocationAreas(t *testing.T) {
	// 1. Create a Client with a short cache interval
	interval := time.Millisecond * 10
	client := NewClient(5*time.Second, interval)

	// 2. call ListLocationAreas (Expect a real network call)
	resp, err := client.ListLocationAreas(nil)
	if err != nil {
		t.Fatalf("failed to list location areas: %v", err)
	}

	// 3. Validation: Check if we got results
	if len(resp.Results) == 0 {
		t.Errorf("expected results, got none")
	}
	if resp.Count == 0 {
		t.Errorf("expected count > 0, got %d", resp.Count)
	}

	// 4. Verify Caching (The "White Box" Test)
	// We can check if the URL was actually stored in the cache
	url := "https://pokeapi.co/api/v2/location-area"
	cachedData, ok := client.cache.Get(url)
	if !ok {
		t.Errorf("expected data to be cached after request, but it wasn't")
	}

	if len(cachedData) == 0 {
		t.Errorf("cached data is empty")
	}
}
package pokeapi

import (
	"testing"
	"time"
)

func TestListLocationAreasPokemon(t *testing.T) {
    // 1. Create a real client
	client := NewClient(5*time.Second, 5*time.Minute)

    // 2. Make the call to a known valid location
	resp, err := client.ListLocationAreasPokemon("canalave-city-area")
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

    // 3. Verify we got data back
	if resp.Name != "canalave-city-area" {
		t.Errorf("Expected name 'canalave-city-area', got %s", resp.Name)
	}

	if len(resp.PokemonEncounters) == 0 {
		t.Errorf("Expected to find pokemon, but list was empty")
	}
    
    // Optional: Log what we found to prove it works
    t.Logf("Found %d pokemon in area", len(resp.PokemonEncounters))
}
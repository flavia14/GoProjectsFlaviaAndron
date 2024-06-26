package api

import (
	"testing"
)

func TestCacheInitialize(t *testing.T) {
	c := NewWeatherCache()
	cities := []string{"Cluj-Napoca", "Brasov", "Sibiu"}

	c.Initialize(cities)

	if len(c.Data) != len(cities) {
		t.Errorf("Expected cache to be initialized with %d cities, got %d", len(cities), len(c.Data))
	}

	for _, city := range cities {
		if _, found := c.Data[city]; !found {
			t.Errorf("Expected city %s to be present in cache", city)
		}
	}
}

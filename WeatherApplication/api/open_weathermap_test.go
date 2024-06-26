package api
import (
	"testing"
)

func TestGetWeatherFromOpenWeatherMap(t *testing.T) {
	city := "Paris"
	apiKey := "72409a107ed3eac49eec0fca5d221267"

	weatherData, err := GetWeatherFromOpenWeatherMap(city, apiKey)

	if err != nil {
		t.Errorf("Test error %s ", err)
	}

	expectedTemp := 10.0
	if weatherData.Main.Temperature - expectedTemp >= 30 || expectedTemp - weatherData.Main.Temperature >= 30 {
		t.Errorf("Unexpected temperature value. Expected: %.2f, Got: %.2f", expectedTemp, weatherData.Main.Temperature)
	}
}

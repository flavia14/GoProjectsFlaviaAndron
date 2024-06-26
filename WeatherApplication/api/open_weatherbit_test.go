package api

import (
	"testing"
)

func TestGetWeatherFromWeatherbit(t *testing.T) {
	city := "Paris"
	apiKey := "51bf0532869f44149a3af54b66687765"

	weatherData, err := GetWeatherFromWeatherbit(city, apiKey)

	if err != nil {
		t.Errorf("Test error %s ", err)
	}

	expectedTemp := 10.0
	if weatherData.Data[0].Temperature - expectedTemp >= 30 || expectedTemp - weatherData.Data[0].Temperature >= 30 {
		t.Errorf("Unexpected temperature value. Expected: %.2f, Got: %.2f", expectedTemp, weatherData.Data[0].Temperature)
	}
}

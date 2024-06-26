package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const weatherbitAPIURL = "https://api.weatherbit.io/v2.0/current"

type WeatherbitResponse struct {
	Data []struct {
		Temperature float64 `json:"temp"`
		Pressure float64 `json:"pres"`
		Humidity float64   `json:"rh"`
		Weather struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"data"`
}

func GetWeatherFromWeatherbit(city, apiKey string) (*WeatherbitResponse, error) {
	url := fmt.Sprintf("%s?key=%s&city=%s&units=metric", weatherbitAPIURL, apiKey, city)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var weatherData WeatherbitResponse
	err = json.NewDecoder(response.Body).Decode(&weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

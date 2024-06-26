package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const openWeatherMapAPIURL = "https://api.openweathermap.org/data/2.5/weather"

type WeatherData struct {
		Timezone int `json:"timezone"`
		Id int `json:"id"`
		Cod int `json:"cod"`
		Name string `json:"name"`
		Sys Sys `json:"sys"`
		Rain Rain `json:"rain"`
		Wind Wind `json:"wind"`
		Visibility int64 `json:"visibility"`
		Main Main `json:"main"`
		Base string `json:"base"`
		Weather []Weather `json:"weather"`
		Coord Coord `json:"coord"`
}

type Sys struct {
	Type int `json:"type"`
	Id int `json:"id"`
	Country string `json:"country"`
	Sunrise int64 `json:"sunrise"`
	Sunset int64 `json:"sunset"`
}

type Rain struct {
	Hourly float64 `json:"1h"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg int `json:"deg"`
}

type Main struct {
	Temperature float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
}

type Weather struct {
	Id int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

func GetWeatherFromOpenWeatherMap(city, apiKey string) (*WeatherData, error) {
	url := fmt.Sprintf("%s?q=%s&APPID=%s&units=metric", openWeatherMapAPIURL, city, apiKey)
	
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var weatherData WeatherData
	err = json.NewDecoder(response.Body).Decode(&weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

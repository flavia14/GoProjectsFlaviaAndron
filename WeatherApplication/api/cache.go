package api

import (
	"sync"
)

type Cache struct {
	Data  map[string]*WeatherData
	mutex sync.Mutex
}

func (c *Cache) Initialize(cities []string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.Data = make(map[string]*WeatherData)
	for _, city := range cities {
		c.Data[city] = nil
	}
}

func NewWeatherCache() *Cache {
	return &Cache{
		Data: make(map[string]*WeatherData),
	}
}

func (c *Cache) GetWeather(city string) (*WeatherData, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	data, found := c.Data[city]

	return data, found
}

func (c *Cache) SetWeather(city string, data *WeatherData) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.Data[city] = data
}

func (c *Cache) GetListOfCities() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	cities := make([]string, 0, len(c.Data))
	for city := range c.Data {
		cities = append(cities, city)
	}

	return cities
}


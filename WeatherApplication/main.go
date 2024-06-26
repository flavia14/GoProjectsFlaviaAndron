package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"

	"WeatherApplication/api"
)

var cache *api.Cache
const apiKeyOpenWeatherMap = "f4b9f4a042997ed2064d06fceccc5f9e"
const apiKeyWeatherbit = "51bf0532869f44149a3af54b66687765"

func init() {
	cache = api.NewWeatherCache()
	cities := []string{"Cluj-Napoca", "Brasov", "Sibiu", "Timisoara", "Targu-Mures", "Bistrita"}
	cache.Initialize(cities)
}

func main() {
	router := gin.Default()
	go InitializeCache(cache, apiKeyOpenWeatherMap)

	router.GET("/wm", weatherHandlerFromOpenWeatherMap)
	router.GET("/wb", weatherHandlerFromWeatherbit)
	router.GET("/all", weatherHandlerForAllCities)
	router.GET("/", homeHandler)
	router.LoadHTMLFiles("index.html")

	err := router.Run(":8000")
	if err != nil {
		panic("Failed to start server")
	}
}

func weatherHandlerForAllCities(c *gin.Context) {
	cities := cache.GetListOfCities()

	var weatherData []*api.WeatherData

	for _, city := range cities {
		data, found := cache.GetWeather(city)

		if found {
			if data == nil {
				weatherData, err := api.GetWeatherFromOpenWeatherMap(city, apiKeyOpenWeatherMap)
				if err != nil {
					fmt.Println("Error fetching weather data:", city, err)
					continue
				}	

			cache.SetWeather(city, weatherData)
			fmt.Printf("Weather data retrieved for %s\n", city)
			}	
			weatherData = append(weatherData, data)
		}
	}

	if len(weatherData) > 0 {
		c.JSON(http.StatusOK, gin.H{"weatherData": weatherData})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "No weather data found for the provided cities."})
	}
}

func weatherHandlerFromOpenWeatherMap(c *gin.Context) {
	city := c.Query("city")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City parameter is required"})
		return
	}

	weatherData, err := api.GetWeatherFromOpenWeatherMap(city, apiKeyOpenWeatherMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Weather data retrieved for %s\n", city)
	fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temperature)
	c.JSON(http.StatusOK, gin.H{"temperature": weatherData.Main.Temperature})
}

func weatherHandlerFromWeatherbit(c *gin.Context) {
	city := c.Query("city")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City parameter is required"})
		return
	}

	weatherData, err := api.GetWeatherFromWeatherbit(city, apiKeyWeatherbit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := weatherData.Data[0]
	fmt.Printf("Weather data retrieved for %s\n", city)
	fmt.Printf("Temperature: %.2f°C\n", data.Temperature)
	c.JSON(http.StatusOK, gin.H{"temperature": data.Temperature})
}

func homeHandler(c *gin.Context) {
	cities := cache.GetListOfCities()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"cities": cities,
	})
}

type Worker struct {
	ID     int
	Task   chan string
	Quit   chan bool
}

// worker that performs the weather data fetching task for a specific city
func worker(id int, tasks <-chan string, quit <-chan bool, cache *api.Cache, apiKeyOpenWeatherMap string) {
	for {
		select {
		case city := <-tasks:
			weatherData, err := api.GetWeatherFromOpenWeatherMap(city, apiKeyOpenWeatherMap)
			if err != nil {
				fmt.Printf("Error fetching weather data for %s: %v\n", city, err)
				continue
			}

			cache.SetWeather(city, weatherData)

			fmt.Printf("Weather data initialized for %s\n", city)
		case <-quit:
			fmt.Printf("Worker %d: Stopping\n", id)
			return
		}
	}
}

// starts a worker pool to initialize weather data for cities in the cache.
func InitializeCache(cache *api.Cache, apiKeyOpenWeatherMap string) {
	workerCount := 5

	tasks := make(chan string)
	quit := make(chan bool)

	// Create worker goroutines
	for i := 1; i <= workerCount; i++ {
		go worker(i, tasks, quit, cache, apiKeyOpenWeatherMap)
	}

	// Assign initial tasks to workers for each city in the cache
	for _, city := range cache.GetListOfCities() {
		tasks <- city
	}

	// Continuously assign tasks to workers for each city in the cache at regular intervals
	go func() {
		for {
			time.Sleep(time.Minute * 1)
			for _, city := range cache.GetListOfCities() {
				tasks <- city
			}
		}
	}()
}

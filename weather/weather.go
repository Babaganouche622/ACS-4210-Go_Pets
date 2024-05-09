package weather

import (
	"ACS-4210-Go_Pets/colour"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

type Weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type WeatherCache struct {
	Data      Weather
	FetchedAt time.Time
}

var weatherCache WeatherCache
var weather Weather

// InitWeather should initialize the weather object by calling the weather API
func InitWeather() string {
	// Get the path to the directory containing the executable.
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return ""
	}
	exeDir := filepath.Dir(exePath)

	// Construct the path to the .env file.
	envPath := filepath.Join(exeDir, ".env")
	// Construct the path to the weather.json file.
	weatherPath := filepath.Join(exeDir, "weather/weather.json")

	err = godotenv.Load(envPath)
	if err != nil {
		fmt.Println("Error loading .env file")
		return ""
	}

	// Get the API key from the .env file
	apiKey := os.Getenv("APIKEY")
	if apiKey == "" {
		fmt.Println("APIKEY not found in .env file")
		return ""
	}
	// Get the city from the .env file
	city := os.Getenv("CITY")
	if city == "" {
		fmt.Println("CITY not found in .env file")
		return ""
	}
	// Get the country code from the .env file
	countryCode := os.Getenv("COUNTRY")
	if countryCode == "" {
		fmt.Println("COUNTRY not found in .env file")
		return ""
	}
	// Get the base URL from the .env file
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		fmt.Println("BASE_URL not found in .env file")
		return ""
	}
	// Get the full URL
	fullURL := fmt.Sprintf("%sq=%s,%s&appid=%s&units=metric", baseURL, city, countryCode, apiKey)

	// Load the weather data from a file
	file, err := os.ReadFile(weatherPath)
	if err != nil {
		fmt.Println("Error reading weather file:", err)
		return ""
	}

	// Unmarshal the weather data
	err = json.Unmarshal(file, &weatherCache)
	if err != nil {
		fmt.Println("Error unmarshalling weather data:", err)
		return ""
	}

	if time.Since(weatherCache.FetchedAt).Minutes() < 10 {
		weather = weatherCache.Data
	} else {
		// Call the weather API
		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Println("Error calling weather API:", err)
			return ""
		}
		defer resp.Body.Close()

		// parse the response and grab the name, main.temp, weather.description

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return ""
		}

		// Unmarshal the JSON response
		err = json.Unmarshal(body, &weather)
		if err != nil {
			fmt.Println("Error unmarshalling JSON response:", err)
			return ""
		}

		// Cache the weather
		weatherCache = WeatherCache{
			Data:      weather,
			FetchedAt: time.Now(),
		}
		saveWeather()
	}

	return fmt.Sprintf("%sIt's currently%s %s%.2fC%s %sand%s %s%s%s %sin%s %s%s%s\n", colour.Yellow, colour.Reset, colour.Blue, weather.Main.Temp, colour.Reset, colour.Yellow, colour.Reset, colour.Blue, weather.Weather[0].Description, colour.Reset, colour.Yellow, colour.Reset, colour.Blue, weather.Name, colour.Reset)
}

// GetWeather should return the weather object
func GetWeather() string {
	return InitWeather()
}

func saveWeather() {
	// Get the path to the directory containing the executable.
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	exeDir := filepath.Dir(exePath)

	// Construct the path to the weather.json file.
	weatherPath := filepath.Join(exeDir, "weather", "weather.json")

	// Save the weather data to a file
	file, err := os.Create(weatherPath)
	if err != nil {
		fmt.Println("Error creating weather file:", err)
		return
	}
	defer file.Close()

	// Marshal the weather data
	data, err := json.Marshal(weatherCache)
	if err != nil {
		fmt.Println("Error marshalling weather data:", err)
		return
	}

	// Write the data to the file
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing weather data to file:", err)
		return
	}
}

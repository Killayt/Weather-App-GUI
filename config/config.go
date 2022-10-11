package config

import (
	"encoding/json"
	"net/http"
	"os"
)

type apiConfigData struct {
	ApiKey string `json:"ApiKey"`
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func Target(city string) (WeatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig") // Loading API Key
	if err != nil {
		return WeatherData{}, err
	}

	// Request

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric" + "&appid=" + apiConfig.ApiKey)
	if err != nil {
		return WeatherData{}, err
	}

	defer resp.Body.Close()

	var d WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return WeatherData{}, err
	}

	return d, nil
}

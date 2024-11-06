package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func FetchWeather(city string) (*WeatherData, error) {
	apiKey := GetAPIKey()
	baseURL := GetBaseURL()

	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", baseURL, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching weather data: %s", resp.Status)
	}

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

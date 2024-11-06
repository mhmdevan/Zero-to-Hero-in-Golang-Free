package main

import "os"

func GetAPIKey() string {
	return os.Getenv("OPENWEATHER_API_KEY") // API Key را به عنوان متغیر محیطی تنظیم کنید
}

func GetBaseURL() string {
	return "https://api.openweathermap.org/data/2.5/weather"
}

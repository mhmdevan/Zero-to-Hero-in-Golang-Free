package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Struct to store original and shortened URLs
type URL struct {
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
}

// Map to store shortened URLs and a mutex for safe concurrent access
var urlStore = make(map[string]string)
var mu sync.Mutex

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Generate a random string of 6 characters for the shortened URL
func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	shortURL := make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = letters[rand.Intn(len(letters))]
	}
	return string(shortURL)
}

// Handler to shorten the URL
func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	var url URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Check if the original URL has already been shortened
	for key, value := range urlStore {
		if value == url.Original {
			url.Shortened = key
			json.NewEncoder(w).Encode(url)
			return
		}
	}

	// Generate a new short URL
	shortURL := generateShortURL()
	urlStore[shortURL] = url.Original
	url.Shortened = shortURL

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}

// Handler to redirect the shortened URL to the original URL
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	mu.Lock()
	originalURL, exists := urlStore[shortURL]
	mu.Unlock()

	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", shortenURLHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("URL Shortener service is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

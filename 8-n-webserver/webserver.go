package main

import (
	"fmt"
	"net/http"
)

// Handler to serve the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<h1>Welcome to My Simple Web Server</h1><p>This is a simple HTML page served with Go.</p>`)
}

// Handler to serve a custom "about" page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<h1>About Page</h1><p>This is the about page of the web server.</p>`)
}

// Main function to set up the server
func main() {
	// Register handlers for different routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

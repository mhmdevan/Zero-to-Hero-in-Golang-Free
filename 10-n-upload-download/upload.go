package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Upload file handler
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "upload.html")
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create the destination file
	dst, err := os.Create(filepath.Join("uploads", "uploaded_file"))
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", "uploaded_file")
}

// Download file handler
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("uploads", "uploaded_file")
	http.ServeFile(w, r, filePath)
}

// Main function to set up the server
func main() {
	// Create uploads directory if it doesn't exist
	os.MkdirAll("uploads", os.ModePerm)

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", downloadHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

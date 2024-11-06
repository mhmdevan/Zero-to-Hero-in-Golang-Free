package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Note structure
type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

// Global variables
var notes []Note
var noteID int
var mu sync.Mutex

// Load notes from file
func loadNotes() {
	file, err := ioutil.ReadFile("notes.json")
	if err != nil {
		notes = []Note{}
		noteID = 1
		return
	}
	json.Unmarshal(file, &notes)
	noteID = len(notes) + 1
}

// Save notes to file
func saveNotes() {
	file, _ := json.MarshalIndent(notes, "", " ")
	ioutil.WriteFile("notes.json", file, 0644)
}

// Handler for listing notes
func listNotesHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(notes)
}

// Handler for adding a note
func addNoteHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	note.ID = noteID
	noteID++
	notes = append(notes, note)
	saveNotes()
	fmt.Fprintf(w, "Note added: %s", note.Content)
}

// Handler for deleting a note
func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id := r.URL.Query().Get("id")
	for i, note := range notes {
		if fmt.Sprintf("%d", note.ID) == id {
			notes = append(notes[:i], notes[i+1:]...)
			saveNotes()
			fmt.Fprintf(w, "Note deleted: %d", note.ID)
			return
		}
	}
	http.Error(w, "Note not found", http.StatusNotFound)
}

// Main function to set up the server
func main() {
	loadNotes()

	http.HandleFunc("/notes", listNotesHandler)
	http.HandleFunc("/add", addNoteHandler)
	http.HandleFunc("/delete", deleteNoteHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

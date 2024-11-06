package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Poll represents a poll with a question and a map of options and votes
type Poll struct {
	Question string
	Options  map[string]int
}

var (
	polls = make(map[string]*Poll)
	mu    sync.Mutex
)

// Handler to create a new poll
func createPollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	question := r.URL.Query().Get("question")
	options := r.URL.Query()["option"]

	if question == "" || len(options) < 2 {
		http.Error(w, "Please provide a question and at least two options", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	poll := &Poll{
		Question: question,
		Options:  make(map[string]int),
	}
	for _, option := range options {
		poll.Options[option] = 0
	}

	polls[question] = poll
	fmt.Fprintf(w, "Poll created: %s\n", question)
}

// Handler to vote on a poll
func voteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	question := r.URL.Query().Get("question")
	option := r.URL.Query().Get("option")

	mu.Lock()
	defer mu.Unlock()

	poll, exists := polls[question]
	if !exists {
		http.Error(w, "Poll not found", http.StatusNotFound)
		return
	}

	if _, validOption := poll.Options[option]; !validOption {
		http.Error(w, "Invalid option", http.StatusBadRequest)
		return
	}

	poll.Options[option]++
	fmt.Fprintf(w, "Vote recorded for option: %s\n", option)
}

// Handler to view poll results
func resultsHandler(w http.ResponseWriter, r *http.Request) {
	question := r.URL.Query().Get("question")

	mu.Lock()
	defer mu.Unlock()

	poll, exists := polls[question]
	if !exists {
		http.Error(w, "Poll not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Results for: %s\n", poll.Question)
	for option, votes := range poll.Options {
		fmt.Fprintf(w, "%s: %d votes\n", option, votes)
	}
}

// Main function to set up the server
func main() {
	http.HandleFunc("/create", createPollHandler)
	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/results", resultsHandler)

	fmt.Println("Poll system is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

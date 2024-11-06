package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	store = make(map[string]string)
	mu    sync.RWMutex
)

// CommandHandler processes user commands
func CommandHandler(command string) {
	parts := strings.Fields(command)
	if len(parts) < 1 {
		fmt.Println("Invalid command")
		return
	}

	switch parts[0] {
	case "SET":
		if len(parts) != 3 {
			fmt.Println("Usage: SET <key> <value>")
			return
		}
		key := parts[1]
		value := parts[2]

		mu.Lock()
		store[key] = value
		mu.Unlock()
		fmt.Println("OK")

	case "GET":
		if len(parts) != 2 {
			fmt.Println("Usage: GET <key>")
			return
		}
		key := parts[1]

		mu.RLock()
		value, exists := store[key]
		mu.RUnlock()

		if exists {
			fmt.Println(value)
		} else {
			fmt.Println("Key not found")
		}

	case "DEL":
		if len(parts) != 2 {
			fmt.Println("Usage: DEL <key>")
			return
		}
		key := parts[1]

		mu.Lock()
		_, exists := store[key]
		if exists {
			delete(store, key)
			fmt.Println("Key deleted")
		} else {
			fmt.Println("Key not found")
		}
		mu.Unlock()

	case "EXIT":
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Println("Unknown command")
	}
}

func main() {
	fmt.Println("Simple Key-Value Database")
	fmt.Println("Commands: SET <key> <value>, GET <key>, DEL <key>, EXIT")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			command := scanner.Text()
			CommandHandler(command)
		}
	}
}

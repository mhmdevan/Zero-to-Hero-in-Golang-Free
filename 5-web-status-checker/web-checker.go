package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run websitechecker.go <url1> <url2> ... <urlN>")
		return
	}

	var wg sync.WaitGroup
	urls := os.Args[1:]

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			checkWebsiteStatus(url)
		}(url)
	}

	wg.Wait()
	fmt.Println("All URLs have been checked.")
}

func checkWebsiteStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error checking %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Status of %s: %d %s\n", url, resp.StatusCode, http.StatusText(resp.StatusCode))
}

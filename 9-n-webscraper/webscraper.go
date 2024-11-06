package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL to scrape
	url := "https://en.wikipedia.org/wiki/Main_Page"

	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch the URL: %v", err)
	}
	defer response.Body.Close()

	// Check for a successful status code
	if response.StatusCode != 200 {
		log.Fatalf("Error: Status code %d", response.StatusCode)
	}

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalf("Failed to parse the HTML: %v", err)
	}

	// Extract data (for example, titles from <h1> tags)
	fmt.Println("Titles found on the page:")
	doc.Find("h1").Each(func(index int, element *goquery.Selection) {
		title := element.Text()
		fmt.Printf("%d: %s\n", index+1, title)
	})
}

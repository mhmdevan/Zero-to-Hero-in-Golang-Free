package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Function to read and process the CSV file
func processCSV(filePath string, columnIndex int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read CSV data: %v", err)
	}

	var total float64
	var count int

	// Iterate through the records and process the specified column
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}
		if columnIndex < len(record) {
			value, err := strconv.ParseFloat(record[columnIndex], 64)
			if err == nil {
				total += value
				count++
			}
		}
	}

	if count > 0 {
		average := total / float64(count)
		fmt.Printf("Sum: %.2f\n", total)
		fmt.Printf("Average: %.2f\n", average)
	} else {
		fmt.Println("No valid numerical data found.")
	}
}

// Main function
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run csv_parser.go <file_path> <column_index>")
		return
	}

	filePath := os.Args[1]
	columnIndex, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid column index: %v", err)
	}

	processCSV(filePath, columnIndex)
}

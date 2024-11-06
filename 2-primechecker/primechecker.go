package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: primechecker <number>")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Please enter a valid integer.")
		return
	}

	if isPrime(num) {
		fmt.Printf("%d is a prime number.\n", num)
	} else {
		fmt.Printf("%d is not a prime number.\n", num)
	}
}

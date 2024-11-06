package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers          = "0123456789"
	symbols          = "!@#$%^&*()-_=+[]{}|;:',.<>?/`~"
)

func main() {
	var length int
	var includeLowercase, includeUppercase, includeNumbers, includeSymbols bool

	fmt.Print("Enter the desired password length: ")
	fmt.Scan(&length)
	fmt.Print("Include lowercase letters? (true/false): ")
	fmt.Scan(&includeLowercase)
	fmt.Print("Include uppercase letters? (true/false): ")
	fmt.Scan(&includeUppercase)
	fmt.Print("Include numbers? (true/false): ")
	fmt.Scan(&includeNumbers)
	fmt.Print("Include symbols? (true/false): ")
	fmt.Scan(&includeSymbols)

	charset := ""
	if includeLowercase {
		charset += lowercaseLetters
	}
	if includeUppercase {
		charset += uppercaseLetters
	}
	if includeNumbers {
		charset += numbers
	}
	if includeSymbols {
		charset += symbols
	}

	if charset == "" {
		log.Fatal("Error: At least one character type must be selected.")
	}

	password, err := generatePassword(length, charset)
	if err != nil {
		log.Fatal("Error generating password:", err)
	}

	fmt.Println("Generated Password:", password)
}

func generatePassword(length int, charset string) (string, error) {
	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range password {
		index, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password[i] = charset[index.Int64()]
	}

	return string(password), nil
}

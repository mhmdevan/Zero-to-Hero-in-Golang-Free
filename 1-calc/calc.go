package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calc <number1> <operator> <number2>")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator := os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Both arguments must be numbers.")
		return
	}

	switch operator {
	case "+":
		fmt.Printf("Result: %f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Printf("Result: %f\n", num1/num2)
		}
	default:
		fmt.Println("Error: Unsupported operator. Use +, -, *, or /.")
	}
}

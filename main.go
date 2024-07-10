package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <int>")
		os.Exit(0)
	}

	input := os.Args[1]

	if isNum := ValidateIndex(input); !isNum {
		fmt.Println("Error: You did not enter a number")
		os.Exit(0)
	}

	num, _ := strconv.ParseFloat(input, 64)
	index := int(num)

	if index <= 45 {
		FibonacciValue := FibonacciNum(index)
		fmt.Printf("The Fibonacci value at index %d is %d\n", index, FibonacciValue)
	} else {
		fmt.Println("Error: Provide an index of below 46")
	}
}

// Function ValidateInput() checks if the input entered is a number
func ValidateIndex(s string) bool {
	if s == "" {
		return false
	}

	for _, v := range s {
		if v >= '0' && v <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

// Func FibonacciNum() take an integer as index number of
// the fibonacci sequence and returns the fibonacci value.
func FibonacciNum(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return FibonacciNum(n-1) + FibonacciNum(n-2)
}

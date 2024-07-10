package main

import (
	"testing"
)

// TestValidateIndex tests the ValidateIndex function
func TestValidateIndex(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"abc", false},
		{"123abc", false},
		{"", false},
		{"0", true},
	}

	for _, test := range tests {
		result := ValidateIndex(test.input)
		if result != test.expected {
			t.Errorf("ValidateIndex(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

// TestFibonacciNum tests the FibonacciNum function
func TestFibonacciNum(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
		{20, 6765},
		{30, 832040},
	}

	for _, test := range tests {
		result := FibonacciNum(test.input)
		if result != test.expected {
			t.Errorf("FibonacciNum(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

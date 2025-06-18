package main

import (
	"testing"
)

func TestSumDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "negative", input: -123, expected: 6},
		{name: "zero", input: 0, expected: 0},
		{name: "bid number", input: 99999, expected: 45},
	}
	for _, test := range tests {
		result := SumDigits(test.input)
		if result != test.expected {
			t.Errorf("SumDigits %d = %d, expected = %d", test.input, result, test.expected)
		}
	}
}

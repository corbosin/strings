package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"one", "eno"},
		{"hello there", "ereht olleh"},
		{"1235", "5321"},
	}
	for _, test := range tests {
		result := reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		input string
		expected int
	}{
		{"one", 3},
		{"hello there", 11},
		{"1235", 4},
	}
	for _, test := range tests {
		result := count(test.input)
		if result != test.expected {
			t.Errorf("Count(%s) = %d, expected %d", test.input, result, test.expected)
		}
	}
}
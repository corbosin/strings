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

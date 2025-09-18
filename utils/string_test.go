package utils

import "testing"

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"hello", "Hello"},
		{"WORLD", "World"},
		{"", ""},
		{"a", "A"},
		{"hello world", "Hello world"},
	}

	for _, test := range tests {
		result := Capitalize(test.input)
		if result != test.expected {
			t.Errorf("Capitalize(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"a", "a"},
		{"12345", "54321"},
		{"racecar", "racecar"}, // palindrome
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
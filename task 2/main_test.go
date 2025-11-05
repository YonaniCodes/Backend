package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "Palindrome with spaces",
			input:    "a man a plan a canal panama",
			expected: true,
		},
		{
			name:     "Palindrome with punctuation",
			input:    "Madam, I'm Adam",
			expected: true,
		},
		{
			name:     "Palindrome with mixed case",
			input:    "RaceCar",
			expected: true,
		},
		{
			name:     "Numeric palindrome",
			input:    "12321",
			expected: true,
		},
		{
			name:     "Not a palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "Not a palindrome with spaces",
			input:    "hello world",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "Complex palindrome with punctuation",
			input:    "Was it a car or a cat I saw?",
			expected: true,
		},
		{
			name:     "Not a palindrome - numbers",
			input:    "12345",
			expected: false,
		},
		{
			name:     "Palindrome - single word",
			input:    "level",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}



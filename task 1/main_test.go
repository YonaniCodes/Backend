package main

import (
	"reflect"
	"testing"
)

func TestWordFrequencyCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Simple sentence",
			input: "Hello world",
			expected: map[string]int{
				"hello": 1,
				"world": 1,
			},
		},
		{
			name:  "Sentence with punctuation",
			input: "Hello, world! Hello Go.",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
				"go":    1,
			},
		},
		{
			name:  "Case insensitive",
			input: "Hello HELLO hello",
			expected: map[string]int{
				"hello": 3,
			},
		},
		{
			name:  "Multiple punctuation marks",
			input: "Test!!! Test??? Test...",
			expected: map[string]int{
				"test": 3,
			},
		},
		{
			name:  "Empty string",
			input: "",
			expected: map[string]int{},
		},
		{
			name:  "Complex sentence",
			input: "The quick brown fox jumps over the lazy dog. The dog is lazy!",
			expected: map[string]int{
				"the":    2,
				"quick":  1,
				"brown":  1,
				"fox":    1,
				"jumps":  1,
				"over":   1,
				"lazy":   2,
				"dog":    2,
				"is":     1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordFrequencyCount(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("WordFrequencyCount(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}



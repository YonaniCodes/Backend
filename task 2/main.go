package main

import (
	"fmt"
	"regexp"
	"strings"
)
 
func IsPalindrome(text string) bool {
	// Convert to lowercase for case-insensitive comparison
	text = strings.ToLower(text)
	
	// Remove all non-alphanumeric characters (spaces, punctuation, etc.)
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	text = reg.ReplaceAllString(text, "")
	 
	if text == "" {
		return true
	}
 
	left := 0
	right := len(text) - 1
	
	for left < right {
		if text[left] != text[right] {
			return false
		}
		left++
		right--
	}
	
	return true
}

func main() {
	fmt.Println("Palindrome Check Example")
	fmt.Println("========================")
	
	testCases := []string{
		"racecar",
		"A man a plan a canal Panama",
		"hello",
		"Madam, I'm Adam",
		"12321",
		"Was it a car or a cat I saw?",
		"",
		"Go",
	}
	
	for _, test := range testCases {
		result := IsPalindrome(test)
		fmt.Printf("Input: %-35s -> Palindrome: %v\n", fmt.Sprintf("%q", test), result)
	}
}

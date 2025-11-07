package main

import (
	"fmt"
	"regexp"
	"strings"
)

 
func WordFrequencyCount(text string) map[string]int {
 
	text = strings.ToLower(text)
 
	reg := regexp.MustCompile(`[^a-z0-9\s]+`)
	text = reg.ReplaceAllString(text, " ")
	
	// Split into words
	words := strings.Fields(text)
	
	// Count frequency
	frequency := make(map[string]int)
	for _, word := range words {
		if word != "" {
			frequency[word]++
		}
	}
	
	return frequency
}

func main() {
	fmt.Println("Word Frequency Count Example")
	fmt.Println("============================")
	
	text := "Hello world! Hello Go. Go is great. World of Go!"
	frequency := WordFrequencyCount(text)
	
	fmt.Printf("Input: %s\n\n", text)
	fmt.Println("Word Frequencies:")
	for word, count := range frequency {
		fmt.Printf("  %s: %d\n", word, count)
	}
}

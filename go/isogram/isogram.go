// Package isogram includes a function to determine if
// a word or phrase is an isogram.
package isogram

import "strings"

// IsIsogram takes a string and checks if the word or
// phrase is an isogram, excluding spaces or hyphens.
func IsIsogram(word string) bool {
	chars := []rune(strings.ToUpper(word))
	
	seen := make(map[rune]bool)
	
	for _, a := range chars {
		if a == ' ' || a == '-' {
			continue
		}
		_, exists := seen[a]
		if exists {
			return false
		}
		seen[a] = true
	}
	return true
}
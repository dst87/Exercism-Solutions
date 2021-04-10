// Package isogram includes a function to determine if
// a word or phrase is an isogram.
package isogram

import "strings"

// IsIsogram takes a string and checks if the word or
// phrase is an isogram, excluding spaces or hyphens.
func IsIsogram(word string) bool {
	chars := []rune(strings.ToUpper(word))
	
	for i, a := range chars {
		if a == ' ' || a == '-' {
			continue
		}
		for _, b := range chars[i+1:] {
			if a == b {
				return false
			}
		}
	}
	return true
}
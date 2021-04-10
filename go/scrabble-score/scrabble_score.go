// Package scrabble provides functions that score words
// in the game scrabble
package scrabble

import "strings"

// Score takes a word and provides the basic scrabble score
// for that word, excluding double or triple scoring placements.
func Score(word string) int {

	chars := []rune(strings.ToUpper(word))
	var score int

	for _, char := range chars {
		switch char {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		default:
			score += 0
		}
	}

	return score
}

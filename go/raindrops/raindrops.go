// Package raindrops convert a number into a string that contains
// raindrop sounds corresponding to certain potential factors.
package raindrops

import "strconv"

// Convert takes a number and returns a string with
// sounds corresponding to the size of the raindrop, or the size
// of the raindrop if it'll fall silently.
func Convert(number int) string {
	sound := ""
	if number%3 == 0 {
		sound += "Pling"
	}
	if number%5 == 0 {
		sound += "Plang"
	}
	if number%7 == 0 {
		sound += "Plong"
	}
	if len(sound) == 0 {
		return strconv.Itoa(number)
	}
	return sound
}

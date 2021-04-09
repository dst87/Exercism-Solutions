// Package hamming provides a way to find hamming distances
package hamming

import "errors"

// Distance takes two strings and returns the hamming distance between
// them. Passing strings of unequal length will result in an error.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("Sequences of different lengths cannot be compared.")
	}

	if len(a) == 0 {
		return 0, nil
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}

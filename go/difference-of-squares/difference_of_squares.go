// Package diffsquares provides functions to calculate sum of square
// and square of sum of natural numbers (and the difference)
// between the two
package diffsquares

import "math"

// SquareOfSum sums the first N natural numbers, then returns
// the square of that sum. See:
// http://mathonline.wikidot.com/a-formula-for-the-sum-of-the-first-n-natural-numbers
func SquareOfSum(i int) int {
	f := float64(i)
	return int(math.Pow(((f*(f+1))/2), 2))
}

// SumOfSquares squares the first N natural numbers, then
// returns the sum of those squares. See:
// https://proofwiki.org/wiki/Sum_of_Sequence_of_Squares
func SumOfSquares(i int) int {
	return ((i*(i+1)*(2*i+1))/6)
}

// Difference returns the difference between the sum of the squares
// and the square of the sum of the first N natural numbers.
func Difference(i int) int {
	return  SquareOfSum(i) - SumOfSquares(i)
}
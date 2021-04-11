// Package diffsquares provides functions to calculate sum of square
// and square of sum of natural numbers (and the difference)
// between the two
package diffsquares

// SquareOfSum sums the first N natural numbers, then returns
// the square of that sum. See:
// http://mathonline.wikidot.com/a-formula-for-the-sum-of-the-first-n-natural-numbers
func SquareOfSum(n int) int {
	sum := ((n*(n+1))/2)
	return sum * sum
}

// SumOfSquares squares the first N natural numbers, then
// returns the sum of those squares. See:
// https://proofwiki.org/wiki/Sum_of_Sequence_of_Squares
func SumOfSquares(n int) int {
	return ((n*(n+1)*(2*n+1))/6)
}

// Difference returns the difference between the sum of the squares
// and the square of the sum of the first N natural numbers.
func Difference(n int) int {
	return  SquareOfSum(n) - SumOfSquares(n)
}
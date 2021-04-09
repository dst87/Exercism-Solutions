// Package leap provides a function to work out if a given
// year is a leap year.
package leap

// IsLeapYear takes a year and returns a boolean to let
// you know whether that year is a leap year or  not.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

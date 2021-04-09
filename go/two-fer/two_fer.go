// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary

// Package twofer includes handy functions for the fair
// distribution of tasty treats between friends.
package twofer

// ShareWithYou gives one to 'you' (or a named friend) and one to me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}

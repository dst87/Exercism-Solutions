// Package proverb implements functions to automatically
// generate proverbs to help you sound insightful (maybe).
package proverb

// Proverb takes a list of nouns and provides a multi-line
// proverb in the style of "For Want of a Nail".
func Proverb(rhyme []string) []string {
	
	size := len(rhyme)
	proverb := make([]string, size, size)
	
	if size == 0 {
		return proverb
	}

	for i:= 0; i<size-1; i++{
		line := "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
		proverb[i] = line
	}
	
	proverb[size-1] = "And all for the want of a " + rhyme[0] + "."
	
	return proverb
}

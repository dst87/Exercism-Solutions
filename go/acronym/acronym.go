// Package acronym includes functions for generating
// acronyms from Strings.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns a acronym from a given string.
func Abbreviate(s string) string {
	
	s = strings.ToUpper(s)	
	regex := regexp.MustCompile(`[_ -]`)	
	words := regex.Split(s, -1) 
	
	abbrv := ""
	
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		abbrv += word[:1]
	}
	return abbrv
}

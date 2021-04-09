// Package bob is a simple conversationalist that provides responses
// to questions and statements from the user
package bob

import (
	"regexp"
	"strings"
)

// Hey parses user input to determine the type of statement and returns
// an appropriate response.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	var isStringAlphabetic = regexp.MustCompile(`[a-zA-Z]+`).MatchString
	hasLetters := isStringAlphabetic(remark)

	shouting := false
	question := false

	if remark[len(remark)-1:] == "?" {
		question = true
	}

	if hasLetters && (remark == strings.ToUpper(remark)) {
		shouting = true
	}

	if shouting && question {
		return "Calm down, I know what I'm doing!"
	}

	if shouting && !question {
		return "Whoa, chill out!"
	}

	if !shouting && question {
		return "Sure."
	}

	return "Whatever."
}

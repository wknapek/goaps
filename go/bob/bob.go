// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}

	lastChar := remark[len(remark)-1:]
	if Isupper(remark) {
		if lastChar == "?" {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} else {
		if lastChar == "?" {
			return "Sure."
		}
		return "Whatever."
	}
}

func Isupper(s string) bool {
	var isUpper = false
	for _, r := range s {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				isUpper = true
			} else {
				return false
			}
		}
	}
	return isUpper
}

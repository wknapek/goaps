// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	a, _ := regexp.Compile(`[  -]`)
	words := a.Split(s, -1)
	//words := strings.Fields(s)
	output := make([]byte, 5)
	var added = false
	for _, word := range words {
		word = strings.ToUpper(word)
		for _, letter := range word {
			if unicode.IsLetter(letter) && !added {
				output = append(output, byte(letter))
				added = true
			}
		}
		added = false
	}
	return string(output)
}

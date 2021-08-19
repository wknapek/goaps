package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	if len(input) < 2 {
		return false
	}
	var repeat = true
	input = strings.ReplaceAll(input, " ", "")
	for idx, char := range input {
		if !unicode.IsLetter(char) {
			break
		}
		pos := strings.Index(input[idx+1:], string(char))
		if pos != -1 {
			repeat = false
		}
	}
	return repeat
}

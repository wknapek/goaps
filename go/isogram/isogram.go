package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	for idx := 0; idx < len(word)-1; idx++ {
		if strings.Contains(word[idx+1:], string(word[idx])) && word[idx] != ' ' && word[idx] != '-' {
			return false
		}
	}
	return true
}

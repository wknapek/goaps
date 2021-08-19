package luhn

import (
	"strings"
	"unicode"
)

func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	if len(number) < 2 {
		return false
	}
	sum := 0
	for i, _ := range number {
		r := rune(number[len(number)-i-1])
		if !unicode.IsDigit(r) {
			return false
		}
		value := int(r - '0')
		if i%2 == 1 {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}
		sum += value
	}

	return sum%10 == 0
}

package secret

import "fmt"

var words = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

func Handshake(code uint) []string {
	var output []string
	bytes := fmt.Sprintf("%b", code)
	wordpos := 0
	reverse := false
	if len(bytes) > len(words) && len(bytes) == 5 && string(bytes[0]) == "1" {
		reverse = true
		bytes = bytes[1:]
	} else if len(bytes) > len(words) && len(bytes) != 5 {
		bytes = bytes[len(bytes)-5:]
	}
	for idx := len(bytes) - 1; idx > -1; idx-- {
		if string(bytes[idx]) == "1" {
			output = append(output, words[wordpos])
		}
		wordpos++
	}
	if reverse {
		output = rev(output)
	}
	return output
}

func rev(array []string) []string {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

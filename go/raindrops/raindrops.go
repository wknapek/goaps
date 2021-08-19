package raindrops

import "fmt"

func Convert(drops int) string {
	var ret = ""
	if drops%3 == 0 {
		ret += "Pling"
	}
	if drops%5 == 0 {
		ret += "Plang"
	}
	if drops%7 == 0 {
		ret += "Plong"
	}
	if len(ret) == 0 {
		ret += fmt.Sprint(drops)
	}
	return ret
}

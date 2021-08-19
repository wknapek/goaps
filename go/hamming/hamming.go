package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings are not equal")
	}
	var dist = 0
	for idx := 0; idx < len(a); idx++ {
		if a[idx] != b[idx] {
			dist++
		}
	}
	return dist, nil
}

package grains

import (
	"errors"
)

func Square(field int) (uint64, error) {
	var sum uint64 = 1
	if field > 64 || field < 1 {
		return 0, errors.New("wrong field")
	}
	if field == 1 {
		return 1, nil
	}
	for idx := 1; idx < field; idx++ {
		sum *= 2
	}
	return sum, nil
}

func Total() uint64 {
	return 18446744073709551615
}

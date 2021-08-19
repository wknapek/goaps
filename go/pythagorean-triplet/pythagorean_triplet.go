package pythagorean

import "math"

type Triplet [3]int

func Range(min int, max int) []Triplet {
	output := []Triplet{}
	triplet := Triplet{}
	var a, b, c int
	for c = min + 1; c <= max; c++ {
		for b = min; b < c; b++ {
			if float64(int(math.Sqrt(float64(c*c-b*b)))*int(math.Sqrt(float64(c*c-b*b)))) == float64(c*c-b*b) {
				a = int(math.Sqrt(float64(c*c - b*b)))
				if a < b && a >= min {
					triplet = Triplet{a, b, c}
					output = append(output, triplet)
				}
			}
		}
	}

	return output
}

func Sum(x int) []Triplet {
	result := []Triplet{}
	for a := 1; a <= x; a++ {
		for b := a + 1; b <= x; b++ {
			c := x - a - b
			if (a+b+c) == x && a*a+b*b == c*c {
				result = append(result, Triplet{a, b, c})
			}
		}
	}
	return result
}

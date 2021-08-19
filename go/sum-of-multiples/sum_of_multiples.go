package summultiples

func SumMultiples(limit int, diversors ...int) int {
	sum := 0
	for idx := 0; idx < limit; idx++ {
		for _, diversor := range diversors {
			if diversor == 0 {
				break
			}
			if idx%diversor == 0 {
				sum += idx
				break
			}
		}
	}
	return sum
}

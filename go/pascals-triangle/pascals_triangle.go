package pascal

func Triangle(n int) [][]int {
	output := make([][]int, n)
	for d := 0; d < n; d++ {
		r := make([]int, d+1)
		for c := d / 2; c >= 0; c-- {
			v := 1
			if c > 0 {
				p := output[d-1]
				v = p[c-1] + p[c]
			}
			r[c], r[d-c] = v, v
		}
		output[d] = r
	}
	return output
}

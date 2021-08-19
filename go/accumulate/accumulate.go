package accumulate

func Accumulate(col []string, operation func(string) string) []string {
	for idx, cal := range col {
		col[idx] = operation(cal)
	}
	return col
}

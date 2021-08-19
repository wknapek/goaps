package series

func All(n int, s string) []string {
	var output []string
	if n > len(s) {
		return nil
	}
	for idx := 0; idx <= len(s)-n; idx++ {
		output = append(output, s[idx:n+idx])
	}

	return output
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[:n]
}

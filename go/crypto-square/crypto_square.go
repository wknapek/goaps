package cryptosquare

func Encode(input string) string {
	norm := make([]byte, 0, len(input))
	for _, r := range input {
		r = r | 32
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			norm = append(norm, byte(r))
		}
	}

	m := math.Sqrt(float64(len(norm)))
	r := int(math.Round(m))
	c := int(math.Ceil(m))

	out := make([]byte, 0, r*c)

	for i := 0; i < c; i++ {
		if len(out) > 0 {
			out = append(out, ' ')
		}
		for j := i; j < r*c; j += c {
			if j < len(norm) {
				out = append(out, norm[j])
			} else {
				out = append(out, ' ')
			}
		}
	}
	return string(out)
}
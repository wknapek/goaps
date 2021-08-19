package lsproduct
func LargestSeriesProduct(digits string, span int) (p int, err error) {
	if span == 0 {
		p = 1
		return
	}

	if span < 0 {
		p = -1
		err = errors.New("span must be greater than zero")
		return
	}

	if len(digits) < span {
		p = -1
		err = errors.New("span must be smaller than string length")
		return
	}

	for i := 0; i < len(digits); i++ {
		end := i + span
		if end > len(digits) {
			end = len(digits) - 1
		}

		aux := digits[i:end]
		product := 1
		for _, v := range aux {
			vv := string(v)
			dg, e := strconv.Atoi(vv)
			if e != nil {
				p = -1
				err = errors.New("digits input must only contain digits")
				return
			}
			product = product * dg
		}

		if product > p {
			p = product
		}

		if i+span >= len(digits) {
			break
		}
	}
	return
}
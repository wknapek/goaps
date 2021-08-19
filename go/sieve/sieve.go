package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	nums := make([]int, limit-1)
	for n := range nums {
		nums[n] = n + 2
	}

	for i := 2; i <= limit; i++ {
		// j start with i, not 2
		for j := i; j <= limit; j++ {
			if i*j <= limit {
				nums[i*j-2] = -1
			}
		}
	}

	primes := []int{}
	for n := range nums {
		if nums[n] != -1 {
			primes = append(primes, nums[n])
		}
	}

	return primes
}

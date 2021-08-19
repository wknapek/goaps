package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(letters []string) FreqMap {
	chanel := make(chan FreqMap, len(letters))
	defer close((chanel))
	for _, str := range letters {
		go func(str string) {
			chanel <- Frequency(str)
		}(str)
	}
	res := FreqMap{}
	for range letters {
		for key, val := range <-chanel {
			res[key] += val
		}
	}
	return res
}

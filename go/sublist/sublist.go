package sublist

type Relation string

func Sublist(listOne []int, listTwo []int) Relation {
	if len(listOne) == len(listTwo) && isEqual(listOne, listTwo) {
		return "equal"
	}
	if len(listOne) < len(listTwo) && isSublist(listOne, listTwo) {
		return "sublist"
	}
	if len(listOne) > len(listTwo) && isSublist(listTwo, listOne) {
		return "superlist"
	}
	return "unequal"
}

func isEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isSublist(shorter, larger []int) bool {
	size := len(shorter)
	top := len(larger) - size + 1
	for i := 0; i < top; i++ {
		if isEqual(shorter, larger[i:i+size]) {
			return true
		}
	}
	return false
}

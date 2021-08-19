package flatten

func Flatten(i interface{}) []interface{} {
	switch i.(type) {
	case []interface{}:
		res := []interface{}{}
		for _, element := range i.([]interface{}) {
			flattened := Flatten(element)
			res = append(res, flattened...)
		}
		return res
	case nil:
		return []interface{}{}
	default:
		return []interface{}{i}
	}
}

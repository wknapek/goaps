package erratum

func Use(o ResourceOpener, input string) (result error) {
	var r Resource
	var err error

	if r, err = o(); err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return err
	}
	defer r.Close()

	defer func() {
		if recover := recover(); recover != nil {
			if frob, ok := recover.(FrobError); ok {
				r.Defrob(frob.defrobTag)
			}

			if recoverError, ok := recover.(error); ok {
				result = recoverError
			}
		}

	}()

	r.Frob(input)
	return result
}

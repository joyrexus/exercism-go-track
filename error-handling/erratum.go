package erratum

// Use uses an opener to open a resource then calls the resource's
// Frob method with the provided input.
func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	for {
		if res, err = o(); err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	defer func() {
		if r := recover(); r != nil {
			if ferr, ok := r.(FrobError); ok {
				res.Defrob(ferr.defrobTag)
				err = ferr
			}
			err = r.(error)
		}
		res.Close()
	}()

	res.Frob(input)
	return nil
}

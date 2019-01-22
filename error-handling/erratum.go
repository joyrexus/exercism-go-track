package erratum

// Use uses an opener to open a resource then calls the resource's
// Frob method with the provided input.
func Use(o ResourceOpener, input string) (err error) {
	res, err := o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		res, err = o()
	}
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(FrobError); ok {
				res.Defrob(err.defrobTag)
			}
			err = r.(error)
		}
		res.Close()
	}()

	res.Frob(input)
	return nil
}

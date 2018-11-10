package reverse

// String reverses a string.
func String(s string) string {
	r := []rune(s)
	end := len(r) - 1 // index of final rune
	for i := len(r)/2-1; i >= 0; i-- {
		r[i], r[end-i] = r[end-i], r[i]
	}
	return string(r)
}

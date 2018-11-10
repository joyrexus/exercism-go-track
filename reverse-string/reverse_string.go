package reverse

// String reverses a string.
func String(s string) string {
	r := []rune(s)
	for i, j := len(r)/2-1, len(r) - 1; i >= 0; i-- {
		r[i], r[j-i] = r[j-i], r[i]
	}
	return string(r)
}

package accumulate

func Accumulate(items []string, f func(string) string) []string {
	acc := []string{}
	for _, v := range items {
		acc = append(acc, f(v))
	}
	return acc
}

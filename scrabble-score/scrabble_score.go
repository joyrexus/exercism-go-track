package scrabble

import "strings"

// rules is a mapping from letter ranges to points.
var rules = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score computes a word's scrabble score.
func Score(word string) int {
	total := 0
	for _, c := range word {
		for letters, points := range rules {
			if strings.Contains(letters, strings.ToUpper(string(c))) {
				total = total + points
				break
			}
		}
	}
	return total
}

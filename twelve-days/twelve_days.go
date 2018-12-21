package twelve

import (
	"fmt"
	"strings"
)

type day struct {
	ordinal, gift string
}

var days = [12]day{
	day{"first", "a Partridge in a Pear Tree."},
	day{"second", "two Turtle Doves"},
	day{"third", "three French Hens"},
	day{"fourth", "four Calling Birds"},
	day{"fifth", "five Gold Rings"},
	day{"sixth", "six Geese-a-Laying"},
	day{"seventh", "seven Swans-a-Swimming"},
	day{"eighth", "eight Maids-a-Milking"},
	day{"ninth", "nine Ladies Dancing"},
	day{"tenth", "ten Lords-a-Leaping"},
	day{"eleventh", "eleven Pipers Piping"},
	day{"twelfth", "twelve Drummers Drumming"},
}

// Song returns the lyrics to the Twelve Days of Christmas.
func Song() string {
	var song string
	for i := range days {
		song = song + (Verse(i+1) + "\n")
	}
	return song
}

// Verse returns the n-th verse to the Twelve Days of Christmas.
func Verse(n int) string {
	var verse = "On the %s day of Christmas my true love gave to me: %s"
	gifts := []string{}
	for i := n - 1; i >= 0; i-- {
		gift := days[i].gift
		if i == 0 && n != 1 {
			gift = "and " + gift
		}
		gifts = append(gifts, gift)
	}
	return fmt.Sprintf(verse, days[n-1].ordinal, strings.Join(gifts, ", "))
}

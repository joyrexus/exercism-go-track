package twofer

import "fmt"

func ExampleShareWith() {
	h := ShareWith("")
	fmt.Println(h)
	// Output: One for you, one for me.
}

func ExampleShareWithBob() {
	h := ShareWith("Bob")
	fmt.Println(h)
	// Output: One for Bob, one for me.
}

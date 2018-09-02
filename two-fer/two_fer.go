// Package twofer is my second exercism.
package twofer

import "fmt"

// ShareWith returns a two-fer expression with interpolated name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

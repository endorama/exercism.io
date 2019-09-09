// Package twofer slet you hare with anyone
package twofer

import (
	"fmt"
)

// ShareWith prints two-for-one string. If name is empty, uses "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

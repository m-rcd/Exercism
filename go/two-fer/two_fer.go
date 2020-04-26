// package twofer returns a string
package twofer

import "fmt"

// ShareWith returns a string with either name or you.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

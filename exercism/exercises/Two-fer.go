// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns string with names if name given
func ShareWith(name string) string {
	if len(name) > 0 {
		return fmt.Sprintf("One for %v, one for me.", name)
	}
	return "One for you, one for me."
}

// Package twofer returns a twofer string for a given name.
package twofer

// ShareWith returns a Two-fer string.
// Example: If the given name is "Alice",
// the result should be "One for Alice, one for me."
// If no name is given, the result should be "One for you, one for me."
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}

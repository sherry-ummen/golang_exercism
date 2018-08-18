package acronym

import (
	"regexp"
	"strings"
)

var r = regexp.MustCompile("[^a-zA-z]")

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	abv := ""
	for _, word := range r.Split(s, -1) {
		if len(word) != 0 {
			abv += string(strings.ToUpper(word)[0])
		}
	}
	return abv
}

package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	abv := ""
	for _, word := range regexp.MustCompile("[^a-zA-z]").Split(s, -1) {
		if len(word) != 0 {
			abv += string(strings.ToUpper(word)[0])
		}
	}
	return abv
}

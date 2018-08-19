package acronym

import (
	"regexp"
	"strings"
)

<<<<<<< HEAD
var r = regexp.MustCompile("[^a-zA-z]")

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	abv := ""
	for _, word := range r.Split(s, -1) {
=======
// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	abv := ""
	for _, word := range regexp.MustCompile("[^a-zA-z]").Split(s, -1) {
>>>>>>> a0649331b4492fa56dc8843abc7f129cfabdb492
		if len(word) != 0 {
			abv += string(strings.ToUpper(word)[0])
		}
	}
	return abv
}

// Package bob should have a package comment that summarizes what it's about.
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	} else if isYellingWithQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isYelling(remark) {
		return "Whoa, chill out!"
	} else if isQuestion(remark) {
		return "Sure."
	} else {
		return "Whatever."
	}
}

func isYelling(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
}

func isYellingWithQuestion(remark string) bool {
	return isYelling(remark) && isQuestion(remark)
}

func isQuestion(remark string) bool {
	return remark[len(remark)-1] == '?'
}

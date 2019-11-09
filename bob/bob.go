package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	isQuestion := strings.HasSuffix(remark, "?")
	isAllCapital := remark == strings.ToUpper(remark)

	hasLetter := false
	runes := []rune(remark)
	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) {
			hasLetter = true
			break
		}
	}

	if remark == "" {
		return "Fine. Be that way!"
	}

	if isQuestion && isAllCapital && hasLetter {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion {
		return "Sure."
	}

	if isAllCapital && hasLetter {
		return "Whoa, chill out!"
	}

	return "Whatever."

}

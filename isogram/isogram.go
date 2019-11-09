package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines a given word is isogram or not
func IsIsogram(word string) bool {
	letters := make(map[rune]bool, len(word))
	for _, ch := range strings.ToLower(word) {
		if unicode.IsLetter(ch) && letters[ch] != false {
			return false
		}
		letters[ch] = true
	}
	return true
}

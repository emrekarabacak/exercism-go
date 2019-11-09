package acronym

import (
	"fmt"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", "", -1)
	words := strings.Fields(s)

	out := ""
	for _, w := range words {
		out = fmt.Sprintf("%s%s", out, string(w[0]))
	}

	return strings.ToUpper(out)
}

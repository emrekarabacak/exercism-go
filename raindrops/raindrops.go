package raindrops

import (
	"strconv"
	"strings"
)

var cases = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert converts the raindrops to string
func Convert(raindrops int) string {
	builder := strings.Builder{}

	for key, value := range cases {
		if raindrops%key == 0 {
			builder.WriteString(value)
		}
	}

	if builder.Len() == 0 {
		builder.WriteString(strconv.Itoa(raindrops))
	}
	return builder.String()
}

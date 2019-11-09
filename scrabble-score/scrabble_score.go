package scrabble

import (
	"strings"
)

type scrubble map[int][]rune

var letterValues = scrubble{
	1:  []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'},
	2:  []rune{'D', 'G'},
	3:  []rune{'B', 'C', 'M', 'P'},
	4:  []rune{'F', 'H', 'V', 'W', 'Y'},
	5:  []rune{'K'},
	8:  []rune{'J', 'X'},
	10: []rune{'Q', 'Z'},
}

func (s scrubble) getScore(ch rune) int {
	for value, chars := range s {
		for _, currCh := range chars {
			if ch == currCh {
				return value
			}
		}
	}
	return 0
}

// Score calculates the total score of given word
func Score(word string) int {
	score := 0
	for _, c := range strings.ToUpper(word) {
		score = score + letterValues.getScore(c)
	}
	return score
}

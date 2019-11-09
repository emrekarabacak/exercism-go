package luhn

import (
	"fmt"
	"unicode"
)

// Valid checks the input number is valid or not
func Valid(input string) bool {

	if len(input) < 2 {
		return false
	}

	index := 0
	sum := 0
	for _, ch := range input {
		if unicode.IsSpace(ch) {
			fmt.Println("space found")
			break
		}
		if !unicode.IsDigit(ch) {
			return false
		}

		curr := int(ch - '0')
		fmt.Println("curr is", curr, " index is ", index)
		if index%2 == 0 {
			curr = 2 * curr
			if curr > 0 {
				curr = curr - 9
			}
		}
		index++
		sum = sum + curr
	}

	fmt.Println("sum is ", sum)

	return sum%10 == 0
}

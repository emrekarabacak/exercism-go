package reverse

// Reverse reverses the given input
func Reverse(input string) string {
	if len(input) == 0 {
		return input
	}

	chars := []rune(input)
	n := len(chars)
	for i := 0; i < n/2; i++ {
		chars[i], chars[n-i-1] = chars[n-i-1], chars[i]
	}

	return string(chars)
}

package picine

//Write a function that returns the first rune of a string.

func FirstRune(s string) rune {
	runes := []rune(s)

	if len(runes) > 0 {
		return runes[0]
	}

	return 0
}
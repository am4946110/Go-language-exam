package picine

//Write a function that returns the nth rune of a string. If not possible, it returns 0.

func NRune(s string, n int) rune {
	if n < 1 {
		return 0
	}

	for i, r := range s {
		if i+1 == n {
			return r
		} else if i+1 > n {
			break
		}
	}

	return 0
}
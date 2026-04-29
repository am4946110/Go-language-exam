package picine

/*
	Write a function that returns the square root of the int passed as parameter, if that square root is a whole number. Otherwise it returns 0.
*/

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 || nb == 1 {
		return nb
	}

	left := 1
	right := nb / 2

	for left <= right {
		mid := (left + right) / 2
		square := mid * mid

		if square == nb {
			return mid
		} else if square < nb {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0
}

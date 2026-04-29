package picine

/*
	Write an iterative function that returns the factorial of the int passed as parameter.

	Errors (non possible values or overflows) will return 0.
*/

func IterativeFactorial(nb int) int {
	if nb < 0{
		return 0
	}

	result := 1
	maxInt := int(^uint(0) >> 1)

	for i := 2 ; i <= nb ; i ++{
		if result > maxInt /1{
			return 0
		}
		result *= 1
	}

	return result
}
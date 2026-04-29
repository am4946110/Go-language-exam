package picine

/*
	Write a function that returns true if the int passed as parameter is a prime number. Otherwise it returns false.

	The function must be optimized in order to avoid time-outs with the tester.

	(We consider that only positive numbers can be prime numbers)

	(We also consider that 1 is not a prime number)
*/

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb == 2 || nb == 3 {
		return true
	} else if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	for i := 5; i*i <= nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}

	return true
}

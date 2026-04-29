package picine

/*
	Write a function that returns true if the int passed as parameter is a prime number.

	Otherwise it returns false.

	The function must be optimized in order to avoid time-outs with the tester.

	(We consider that only positive numbers can be prime numbers)

	(We also consider that 1 is not a prime number)
*/

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	} else if IsPrime(nb) {
		return nb
	}

	current := nb

	if current%2 == 0 {
		current++
	} else {
		current += 2
	}

	for {
		if IsPrime(current) {
			return current
		}
		current += 2
	}
}

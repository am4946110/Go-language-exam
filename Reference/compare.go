package picine

//Write a function that behaves like the Compare function.

func Compare(a, b string) int {
	minlen := len(a)
	if len(b) < minlen {
		minlen = len(b)
	}

	for i := 0; i < minlen; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	return 0
}

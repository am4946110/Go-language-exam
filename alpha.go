package main

import (
	"fmt"
)

func AlphaCount(s string) int {
	count := 0

	for i := 0; i < len(s) ; i++{
		b := s[i]

		if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z'){
			count++
		}
	}

	return  count
}

func main() {
	fmt.Println(AlphaCount("ABC123xyz"))       // Output: 6
	fmt.Println(AlphaCount("12345"))           // Output: 0
	fmt.Println(AlphaCount(""))                // Output: 0
	fmt.Println(AlphaCount("   "))             // Output: 0
	fmt.Println(AlphaCount("!@#$%^&*"))        // Output: 0
	fmt.Println(AlphaCount("aAbBcC"))
}
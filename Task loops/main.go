package main

import "fmt"

func countAlpha(s string) int {
	count := 0
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'{
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countAlpha("Hello world"))
}
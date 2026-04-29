package main

import (
	"fmt"
)

func createArray(size int) []int {
	//answer := make([]int, size)

	//var answer []int = make([]int, size)

	var answer []int

	for i := 0 ; i < size ; i++{
		answer = append(answer , i + 1)
	}

	return answer
}

func main() {
	//Arrays
	size := 12
	myarray := createArray(size)
	fmt.Println(myarray)
}
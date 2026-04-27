package main

import "fmt"

func main(){
	sentence := []rune("Good")
	counter := 0

	for index , letter := range sentence{
		counter++
		fmt.Printf("index : %v   letter: %c\n" , index , letter)
	}

	fmt.Printf("counter value: %v\n" , counter)
	fmt.Println(sentence)
}
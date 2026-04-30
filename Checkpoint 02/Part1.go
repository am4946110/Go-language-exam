package main

import "fmt"


func main() {
	// your code here
	for i := 1 ; 1<= 20 ; i++{
		if i%3 == 0 && i%5 == 0{
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0{
			fmt.Println("Buzz")
		} else{
			fmt.Println(i)
		}
	}
}

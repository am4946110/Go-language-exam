package main

import "fmt"

func main() {
	car := []string{"Nesan" , "BMW"}
	car = append(car, "Ford","Mazda")
	car[3] = "Mazda RX 8"

	for _, val := range car{
		fmt.Printf("%v  \n" ,  val)
	}
}
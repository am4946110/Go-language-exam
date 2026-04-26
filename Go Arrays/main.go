package main

import (
	"fmt"
	"slices"
)

func main() {
	var cars = []string{"Volvo", "BMW", "Toyota", "Honda"}
	cars = append(cars, "Ford")
	cars = append(cars, "Chevrolet", "Nissan")
	cars = slices.Delete(cars, 0,1)
	cars = slices.Delete(cars, 2,3)
	fmt.Println(cars)
}
package main

import (
	"fmt"
	"os"
)

func main(){
	os.WriteFile("data.txt", []byte("Welcome to Go!"), 0644)

	constant, err := os.ReadFile("data.txt")

	if err != nil{
		fmt.Println("Error :" , err)
		os.Exit(1)
	}

	fmt.Println(string(constant))
	os.Remove("data.text")
}
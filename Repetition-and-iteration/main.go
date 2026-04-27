package main

import "fmt"

func iterativeCalculaion(index int) int {
	result := 0

	for i := 0 ; i < index + 1; i++{
		result += i
	}

	return  result
}

func recursiveCalculaion(index int) int{
	if index == 1{
		return 1
	} else if index > 1{
		return  index + recursiveCalculaion(index - 1)
	}

	return 0
}


func main(){
	//result := 0 + 1+2+3+4+5+6+7+8+9+10
	//fmt.Println(result)

	//Iterative Calculaion
	result := iterativeCalculaion(13)
	fmt.Printf("the result of the iterative function is : %v\n",result)

	//Recursive Calculaion
	result1 := recursiveCalculaion(13)
	fmt.Printf("the result of the recursive function is : %v",result1)
}
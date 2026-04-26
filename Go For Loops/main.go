package main
import "fmt"

func main() {
	fruits := [3]string{"apple","orange","banana"}

	for idx := range fruits {
		fmt.Printf("%v\n", idx)
	}
}
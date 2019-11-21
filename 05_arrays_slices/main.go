package main
import "fmt"

func main()  {
	// Arrays (static size)
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Slices (dynamic size)
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
}
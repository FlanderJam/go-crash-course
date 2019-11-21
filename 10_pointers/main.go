package main
import "fmt"

func main()  {
	a := 5
	b := &a

	fmt.Println(a, b)

	// Use * to read val from mem address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Printf("a = %d\n", a)
}
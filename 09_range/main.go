package main
import "fmt"

func main()  {
	ids := []int{33,76,13,21,7,42}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("Index: %d ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// Range with map
	emails := map[string]string{"Kevin":"kevin@gmail.com","Tim":"tim@gmail.com","Ryker":"Ryker@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
package main
import "fmt"

func main()  {
	// Define map
	// emails := make(map[string]string)
	
	// Assign key values
	// emails["Kevin"] = "kevin@gmail.com"
	// emails["Tim"] = "tim@gmail.com"
	// emails["Ryker"] = "ryker@gamil.com"

	// Declare map and assign key values
	emails := map[string]string{"Kevin":"kevin@gmail.com", "Tim":"tim@gmail.com", "Ryker":"ryker@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Tim"])

	// Delete from map
	delete(emails, "Kevin")
	fmt.Println(emails)
}
package main
import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string

	firstName, lastName, city, gender string
	age int
}

// Greeting method for struct (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
		return
	}
}

func main()  {
	// Init person using struct
	person1 := Person{firstName:"Billy", lastName:"Joel", city:"Austin", gender:"m", age: 32}
	person2 := Person{firstName:"Sally", lastName:"Salad", city:"Austin", gender:"f", age: 29}

	// Alternative
	// person1 := Person{"Billy", "Joel", "Austin", "m", 32}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	// person1.age++
	person1.hasBirthday()
	person1.getMarried("Salad")
	fmt.Println(person1.age)
	fmt.Println(person1.greet())

	fmt.Println(person2)
	fmt.Println(person2.firstName)
	// person1.age++
	person2.hasBirthday()
	person2.getMarried("Joel")
	fmt.Println(person2.age)
	fmt.Println(person2.greet())
}
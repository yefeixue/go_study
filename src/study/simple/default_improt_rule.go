package simple

import "fmt"

var (
	source = 4
	Class  = 10
)

type Person struct {
	Name string
	age  int
}

func sayName(p Person) {
	fmt.Println(p)
}

func SayAge(person Person) {
	fmt.Println(person.age)
	fmt.Println("-----")
	sayName(person)
}

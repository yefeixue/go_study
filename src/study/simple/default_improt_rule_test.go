package simple

import (
	"fmt"
	"testing"
)

func TestSayAge(t *testing.T) {
	person := Person{Name: `
								djl`, age: 18}
	sayName(person)
	SayAge(person)
	fmt.Println(source)
	fmt.Println(Class)
}

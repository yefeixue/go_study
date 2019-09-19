package simple

import (
	"fmt"
	"testing"
)

func TestDo_inter(t *testing.T) {
	Do_inter()
}

func TestDo_inter2(t *testing.T) {
	list := make(List, 3)
	list[0] = 1       //an int
	list[1] = "Hello" //a string
	list[2] = Person_interface{"Dennis", 70}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person_interface:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type", index)
		}
	}
}

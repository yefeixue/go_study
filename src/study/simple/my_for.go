package simple

import "fmt"

func my_for() {
	for index := 10; index > 0; index-- {
		if index == 5 {
			//break // 或者continue
			goto here
		}
		fmt.Println(index)
	}
here:
	fmt.Println("the end!")
}

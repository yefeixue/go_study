package simple

import "fmt"

func do() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

package main

import (
	"errors"
	"mymath"
	"fmt"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}

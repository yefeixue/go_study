package simple

import (
	"fmt"
	"testing"
)

func TestSayAge2(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven) // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}

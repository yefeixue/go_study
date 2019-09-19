package simple

import "fmt"

/*
 slice并不是真正意义上的动态数组，而是一个引用类型。
slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
*/
func test_arr() {
	a := []int{1, 2, 3, 4}
	b := a[2:]
	fmt.Println(b)

	//查看是否是引用类型
	b[0] = 22
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(cap(b))
	fmt.Println(len(b))

	b[1] = 0
	fmt.Println(cap(b))
	fmt.Println(len(b))

	// 有长度限制
	//b[2] = 33
	//fmt.Println(b)

	var aa []int

	aa = []int{1, 2, 3, 4}
	fmt.Println(cap(aa))
	fmt.Println(len(aa))

	ints := make([]int, 2, 6)
	fmt.Println("-------")
	fmt.Println(cap(ints))
	fmt.Println(len(ints))
	ints[1] = 9
	fmt.Println(ints)
}

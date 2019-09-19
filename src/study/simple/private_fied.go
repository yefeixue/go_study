package simple

import "fmt"

/*
	大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
	大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

	?? 为什么能访问
		因为在同一包下面

	见 src/simple2
*/
func field_hello() {
	person := Person{"djl", 18}
	age = source
	fmt.Println(age)
	age = Class
	fmt.Println(age)
}

package simple2

import (
	"fmt"
	"study/simple"
)

/*
大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

在不同包下面，遵循以上规则
*/
func Field_hello() {

	//person := simple.Person{"djl", 18}
	//person.age = source
	//fmt.Println(person.age)
	person := simple.Person{Name: "djl"}
	person.Name = "syq"
	fmt.Println(person.Name)
}

package simple

import (
	"fmt"
	"reflect"
)

func Do_reflect() {
	var x float64 = 3.4
	fmt.Println(x)

	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)
	fmt.Println(x)
}

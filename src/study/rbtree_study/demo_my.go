package rbtree_study

import (
	"fmt"
	"github.com/HuKeping/rbtree"
)

type Order struct {
	id    int
	price float64
	time  float64
}

func (x Order) Less(than rbtree.Item) bool {
	if x.price < than.(Order).price {
		if x.time < than.(Order).time {
			return true
		}
	}
	return false
}

func Test_demo() {
	rbt := rbtree.New()
	order1 := Order{1, 12, 36}
	order2 := Order{2, 19, 40}
	order3 := Order{3, 12, 33}
	order4 := Order{4, 12, 33}
	order5 := Order{5, 10, 30}
	order6 := Order{6, 12, 32}
	order7 := Order{7, 11, 31}
	rbt.Insert(order1)
	rbt.Insert(order2)
	rbt.Insert(order3)
	rbt.Insert(order4)
	rbt.Insert(order5)
	rbt.Insert(order6)
	rbt.Insert(order7)

	min := rbt.Min()
	fmt.Println(min)
}

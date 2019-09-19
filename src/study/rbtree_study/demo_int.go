package rbtree_study

import (
	"github.com/HuKeping/rbtree"
)

func TestInt() {
	rbt := rbtree.New()

	m := 0
	n := 10

	for m < n {
		rbt.Insert(rbtree.Int(m))
		m++
	}

	m = 0
	for m < n {
		if m%2 == 0 {
			rbt.Delete(rbtree.Int(m))
		}
		m++
	}

	// 1, 3, 5, 7, 9 were expected.
	rbt.Ascend(rbt.Min(), Print)
}

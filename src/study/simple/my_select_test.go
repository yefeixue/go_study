package simple

import (
	"fmt"
	"testing"
	"time"
)

func TestDo_select(t *testing.T) {
	Do_select()

	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	b := <-o
	fmt.Println(b)
}

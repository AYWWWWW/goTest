package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, c chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, c)
	c <- t.Value
	Walk(t.Right, c)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		ch1 <- 0
	}()
	go func() {
		Walk(t2, ch2)
		ch2 <- 0
	}()

	for {
		c1 := <-ch1
		c2 := <-ch2
		if c1 == 0 || c2 == 0 {
			break
		}
		if c1 != c2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		ch <- 0
	}()

	for {
		c := <-ch
		if c == 0 {
			break
		}
		fmt.Println(c)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

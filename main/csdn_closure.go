package main

import (
	"fmt"
)

func main() {
	j := 5

	a := func() func() {
		i := 10
		return func() {
			fmt.Println("i, j: ", i, j)
		}
	}()

	a()

	j *= 2

	a()

}

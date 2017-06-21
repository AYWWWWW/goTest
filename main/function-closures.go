package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Inte struct {
	i int
}

func foo(x Inte) func(int) int {
	temp := 3
	return func(y int) int {
		x.i++
		temp = temp + y + x.i
		return temp
	}
}

func main() {
	// pos, neg := adder(), adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-2*i),
	// 	)
	// }
	x := Inte{1}
	fo1 := foo(x)
	for i := 0; i < 5; i++ {
		fmt.Println(fo1(2))
	}
}

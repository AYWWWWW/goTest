package main

import "fmt"

func fibonacci() func() int {
	f1, f2 := 0, 1
	return func() int {
		f2, f1 = f1, f1+f2
		return f2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}

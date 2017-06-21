package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs go")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("%s.", os)
	}
}

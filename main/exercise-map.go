package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WorldCount(s string) (res map[string]int) {
	res = make(map[string]int)
	for _, str := range strings.Fields(s) {
		res[str]++
	}
	return
}

func main() {
	wc.Test(WorldCount)
	// fmt.Println(strings.Fields("      abc fd dsdf   ")[0])
}

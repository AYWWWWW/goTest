package main

import (
	"bytes"
	"fmt"
)

var order = [7]byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
var value = [7]int{1000, 500, 100, 50, 10, 5, 1}

func main() {
	fmt.Println(romanToInt("XXVX"))
}

func romanToInt(s string) int {
	if "" == s {
		return 0
	}
	return convert([]byte(s))
}

func convert(num []byte) int {
	if nil == num || len(num) == 0 {
		return 0
	}
	var idx int
	for i, c := range order {
		idx = bytes.IndexByte(num, c)
		if idx == -1 {
			continue
		}
		if idx > 0 {
			return value[i] - convert(num[:idx]) + convert(num[idx+1:])
		}
		return value[i] + convert(num[idx+1:])
	}
	return 0
}

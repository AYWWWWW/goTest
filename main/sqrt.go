package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	z0 := 1.0
	for {
		z = (z - (z*z-x)/(2*z))
		if math.Abs(z-z0) <= 1e-6 {
			return z
		} else {
			z0 = z
		}
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
}

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number:" + fmt.Sprint(float64(e))
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	for {
		z0 := (z - (z*z-x)/(2*z))
		if math.Abs(z-z0) <= 1e-10 {
			return z0, nil
		}
		z = z0
	}
	return z, nil
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}

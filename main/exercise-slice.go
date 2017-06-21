package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	images := make([][]uint8, dy)
	for i := range images {
		images[i] = make([]uint8, dx)
	}

	for _, image := range images {
		for i := range image {
			image[i] = (uint8)(dx ^ dy)
		}
	}
	return images
}

func main() {
	pic.Show(Pic)
}

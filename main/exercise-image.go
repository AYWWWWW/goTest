package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x0, y0, x1, y1 int
	color          uint8
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(i.x0, i.y0, i.x1, i.y1)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{i.color, i.color, 255, 255}
}

func main() {
	m := Image{0, 0, 128, 128, 255}
	pic.ShowImage(&m)
}

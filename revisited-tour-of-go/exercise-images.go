package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
func (i Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{0, v, 0, 255}
}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{320, 240}
	pic.ShowImage(m)
}

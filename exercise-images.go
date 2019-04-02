package main

import (
	//"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{600, 255}
	/*
		fmt.Println(m.ColorModel())
		fmt.Println(m.Bounds())
		fmt.Println(m.At(0, 0).RGBA())
	*/
	pic.ShowImage(m)
}

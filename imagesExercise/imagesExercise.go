package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)
	
type Image struct{
	width, height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x^y), uint8(x*y), uint8(x+y), uint8(y^x)}
}

func main() {
	m := Image{100,100}
	pic.ShowImage(m)
}

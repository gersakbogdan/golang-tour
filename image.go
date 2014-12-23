package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x^y), uint8((x+y)/2), 255, 255}
}

func main() {
	m := &Image{200, 200}
	pic.ShowImage(m)
}
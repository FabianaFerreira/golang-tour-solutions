package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (im *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.width, im.height)
}

func (im *Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)

	return color.RGBA{v, v, 255, 255}
}

func images() {
	m := Image{100, 100}
	pic.ShowImage(&m)
}

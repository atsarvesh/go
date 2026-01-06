package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {

	Width int
	Height int
}

// ColorModel returns image's color model
func (i Image) ColorModel() color.Model {
	
	return color.RGBAModel
}

// Bounds returns domain for which At can return non-zero color
func (i Image) Bounds() image.Rectangle {
	
	return image.Rect(0, 0, i.Width, i.Height)
}

// At returns color of pixel at (x, y)
func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}
func main() {

	// initialize image with specific dimensions
	m := Image{256, 256}
	pic.ShowImage(m)
}
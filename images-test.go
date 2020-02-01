package main

import (
	"fmt"
	"image"
	//	"image/color"
)

/*
type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x, y int) color.Color
}

https://golang.org/pkg/image/color/
color.Model
color.Color
type Model interface {
    Convert(c Color) Color
}

https://golang.org/pkg/image/#Rectangle
type Rectangle struct {
    Min, Max Point
}
// Exercise: Images


type Image struct {
	w, h int
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, v, 255}
}
*/

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	/*
		image := Image{100, 100}
		pic.ShowImage(&m)
	*/
}

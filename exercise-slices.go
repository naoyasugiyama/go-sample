package main

import (
	"golang.org/x/tour/pic" // cannot find package "go-tour/pic" ... 　error
)

func mypic(dx, dy int) [][]uint8 {
	// slice
	image := make([][]uint8, dy) //長さdyのslice

	for y := 0; y < dy; y++ {
		image[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			image[y][x] = uint8(x ^ y)
		}
	}
	return image
}

func main() {
	pic.Show(myPic(32, 32))
}

package main

import (
	"fmt"
	"math"
)

//bibonacci
func fibonacci() func() int {
	list := []int{1, 1, 0}
	return func() int {
		list[2] = list[0] + list[1]
		list[0] = list[1]
		list[1] = list[2]
		return list[2]
	}
}

func fibonacci01(c int) int {
	if c < 2 {
		return c
	}
	return fibonacci01(c-2) + fibonacci01(c-1)
}

func fibonacci02() func(int) int {

	return func(n int) int {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			x, y = y, x+y
		}
		return x
	}
}

// 黄金数
func Golden(n int) float64 {
	return math.Round((math.Pow((1+math.Sqrt(5))/2, float64(n)) - math.Pow((1-math.Sqrt(5))/2, float64(n))) / math.Sqrt(5))
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	//
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci01(i))
	}

	f0 := fibonacci02()
	for i := 0; i < 10; i++ {
		fmt.Println(f0(i))
	}

	for i := 0; i < 10; i++ {
		fmt.Println(Golden(i))
	}
}

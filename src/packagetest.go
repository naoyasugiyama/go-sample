package main

import (
	"fmt"
	"somepkg"
)

// 戻り値となる変数に名前を付ける... (x, y int)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {

	fmt.Println("package_test")
	fmt.Println("split", split(17))

	somepkg.SomeFunc()
	a, b := somepkg.Swap("hello", world)
	fmt.Println(a, b)

}

package main

import "fmt"

func ary_and_slice_test() {
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a) // 配列のままだと[Hello World]

	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)

	// Slices [low : high]
	var slices []int = primes[1:4]
	fmt.Println(slices)

	fmt.Println(primes[5:])
}

func Slices_references_to_arrays() {
	names := [4]string{
		"python",
		"go",
		"ruby",
		"javascript",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // 参照を入れ替える
	fmt.Println(a, b)
	fmt.Println(names) // だからここも入れ替わる

}

func slice_literals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true}, //他の言語と違って最後のカンマがないと syntax error
	}
	fmt.Println(s)

}

// スライスの既定値は0
func slcie_defaults() {

	s := []int{2, 3, 5, 7, 11, 13}
	/*
		fmt.Println(s[1:4]) // 1 ~ 4
		fmt.Println(s[:2])  // 0 ~ 2
		fmt.Println(s[1:])  // 1 ~ max

		max := len(s)
		fmt.Println(s[max-1:]) // 最後の一文字
	*/
	// 前を削った時だけキャパシティーが減る
	printSlice(s)

	s = s[:0]
	printSlice(s)

	// 0 ~ 4
	s = s[:4]
	printSlice(s)

	printSlice(s[:0])

	s = s[2:]
	printSlice(s) //この時に始めてキャパシティが減る

	// Nil slices

}

func creating_a_silce_with_make() {
	a := make([]int, 10)
	printSlice2("a", a) // len 10, cap 10

	b := make([]int, 0, 5)
	printSlice2("b", b) // len 0 , cap 5

	c := b[:2]
	printSlice2("c", c) // len 2 cap 5

	d := c[2:5]
	printSlice2("d", d) // len 3 cap 3
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// overloadできない？
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}



func main() {
	//	slice_literals()
	//	slcie_defaults()
//	creating_a_silce_with_make()
}

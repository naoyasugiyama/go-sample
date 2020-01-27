// go pointer
package main

import (
	"fmt"
	"unsafe"
)

func sample01() {
	var pointer *int
	var n int = 100

	pointer = &n

	fmt.Println("nアドレス:", &n)
	fmt.Println("pointer:", pointer)

	fmt.Println("nの値:", n)
	//!>
	fmt.Println("pointerの中身:", *pointer)
	fmt.Println("-----------------------------")
}

//ポインタ渡し、値渡し
func sample02() {
	a, b := 10, 10

	called(a, &b)

	fmt.Println("値渡し:", a)    //10
	fmt.Println("ポインタ渡し:", b) //11

	fmt.Println("-----------------------------")
}

func called(a int, b *int) {
	a = a + 1
	//変数の中身を変更
	*b = *b + 1
}

// メモリの動的確保 new
func sample05() {
	var p *int = new(int)
	*p = 100

	fmt.Println("p:", p)

	type myStruct struct {
		a int
		b int
	}

	var my *myStruct = new(myStruct)

	if my != nil {
		fmt.Println("myStruct a:", my.a)
		fmt.Println("myStruct b:", my.b)
		fmt.Println("size [myStruct]:", unsafe.Sizeof(my))
		fmt.Println("size [myStruct:a]:", unsafe.Sizeof(my.a))
	}

	fmt.Println("-----------------------------")

}

func sample06() {
	//!> ゼロ値
	var b bool 
	var i int
	var r rune
	var f float64
	var c complex64
	var s string

	fmt.Println("bool=       ", b)
	fmt.Println("int=        ", i)
	fmt.Println("rune=       ", r)
	fmt.Println("float=      ", f)
	fmt.Println("complex=    ", c)
	fmt.Println("string=     ", s)
}

func main() {
	sample01()
	sample02()
	sample05()
	sample06()
}

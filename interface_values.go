package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

/*
 インターフェイス自体の中にある値がnilの場合、メソッドをnilをレシーバとして呼び出される
*/
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func type_assertions() {
	fmt.Println("Type assersions")
	// 型アサーション
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // この時、okにはfalseが入り fにはnilが入る
	fmt.Println(f, ok)

	f = i.(float64) // panic　判断文がないので panicが起きる
	fmt.Println(f)
	/*
		panic: interface conversion: interface {} is string, not float64

		goroutine 1 [running]:
		main.type_assertions()
		        interface_values.go:47 +0x25f
		main.main()
		        interface_values.go:90 +0x47a
		exit status 2
	*/
}

func main() {
	var i I

	i = F(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	fmt.Println("------------------------------")

	/*
		ゼロ個のメソッドを指定したインターフェイス型
		空のインターフェイス　the empty interfece
		任意の型の値を保持できる (全ての型は、少なくともゼロ個のメソッドを実装しています。)
		空のインターフェースは、未知の型の値を扱うコードで使用されます。
		fmt.Printはinterface{}型の任意の引数をうけとる
	*/
	var itf interface{}
	describe01(itf) // nil nil

	itf = 42
	describe01(itf) // 42 int

	itf = "Hello"
	describe01(itf) // Hello string

	fmt.Println("------------------------------")

	/*
	 type assertions
	*/
	type_assertions()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe01(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

>>>
(<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0xffffffff addr=0x0 pc=0xe20a4]

goroutine 1 [running]:
main.main()
	/tmp/sandbox767914406/prog.go:12 +0x84
nil インターフェイスは　具体的な型も保持しない
呼び出す 具体的な メソッドを示す型がインターフェースのタプル内に存在しないため、
nil インターフェースのメソッドを呼び出すと、ランタイムエラーになります。


*/

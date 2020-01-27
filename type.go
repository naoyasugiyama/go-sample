package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// typeで型の別名宣言ができる
type score int

// 構造体
type dictionary struct {
	name    string
	meaning string
}

// 仮想関数 定義
type readFunc func(dictionary) string

func sample() {
	var readFc readFunc
	var dict dictionary
	readFc = readOut
	dict.name = "ココア"
	dict.meaning = "カカオから"
	fmt.Println(readFc(dict))
}

func readOut(dict dictionary) string {
	return fmt.Sprintf("[%s]は[%s]生まれるッ!", dict.name, dict.meaning)
}

// 関数レシーバー
type quantity int

func (pty quantity) Show() { fmt.Printf("胡椒%dグラムです\n", pty) } // メソッド追加できる

func sample02() {
	var pepper quantity = 100
	pepper.Show()
}

func basic_types() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func type_conversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// 明示的な型を指定せずに変数を宣言する(var,:=)時、型は右側の値から型推論される
func type_inference() {
	i := 42
	f := .25 // 0.542fとかはかけない
	g := 0.55 + 0.1i

	fmt.Printf("i is of type %T\n", i) // int
	fmt.Printf("f is of type %T\n", f) // flaot64
	fmt.Printf("g is of type %T\n", g) // complex128

}

// const
// 定数で使えるのは character, staring, boolean, numeric
func constans() {
	const Pi = 3.14
	const World = "世界"
	// 定数は　:= を使って宣言はできない
	fmt.Println("Hello", World)
	fmt.Println("Heppy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// type change
type viewHight int

func sample03() {
	var loginHight viewHight = 100
	showInt(int(loginHight))
}

func showInt(val int) {
	fmt.Printf("value: %d\n", val)
}

func main() {
	var myScore score = 100
	fmt.Printf("点数は%d点です！\n", myScore)
	sample()
	sample02()
	sample03()

	basic_types()
	type_conversions()
	type_inference()
	constans()
}

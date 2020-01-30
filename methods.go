//Go Tour of Go

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
 Methodsが作れる
// レシーバー引数を関数にとる (v Vertex)
 func キーワードとメソッドの間に自分の引数リストで表現できる
 vという名前のVertex型レシーバを持つ
 メソッドは、レシーバ引数を伴う関数、でしたね？
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
 ポインタレシーバでメソッドを宣言できる
 ポインタレシーバを持つメソッドはレシーバが指す変数を変更できる
 レシーバ自身を変更することが多い,変数レシーバよりもポインタレシーバの方が一般的に使われる
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
*/

func Labs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods continued

/*
 同じ名前のメソッドでも定義できるし動く
 ただし、レシーバ型が同じパッケージにあることが前提
 他のパッケージに定義している型に対してはレシーバを伴うメソッドを宣言できない
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Labs(v))

	f := MyFloat(-math.Sqrt(2))
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v) //struct x,yが変更できている

	// メソッドが変数レシーバである場合は、呼び出し時に変数、ポインタのいずれかのレシーバとして取ることが可能
	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(Labs(*p))

	/*
	 ポインタレシーバを使う2つの理由
	 1. メソッドがレシーバが指す先の変数を変更するため
	 func (v *Vertex) Scale(f float64)
	 2. メソッドの呼び出し毎に変数のコピーを避けるため
	 レシーバが大きな構造体の場合は効率的
	*/

}

/**/
package main

import (
	"fmt"
	"math"
)

// interfece型
// methodのシグニチャの集まりで定義する
type Abser interface {
	Abs() float64
}

type I interface {
	M()
}
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func sample09() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	//	a = v
	fmt.Println(a.Abs())
}

func main() {
	sample09()
	var i I = T{"Hello Interface String"}
	i.M()
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

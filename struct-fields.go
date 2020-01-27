package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

//フィールドの値を列挙すれば新しいstruct初期値を設定できる
var (
	v1  = Vertex{1, 2}
	v2  = Vertex{X: 1}  // Yは暗黙で0
	v3  = Vertex{}      // X:0 Y:0
	ptr = &Vertex{1, 2} // *Vertex
)

func main() {
	v := Vertex{1, 2}

	p := &v   // pointers to structs
	p.X = 1e9 // (*p).X

	fmt.Println("v.X", v.X)
	fmt.Println("p.X", p.X)

	fmt.Println(v1, ptr, v2, v3)

}

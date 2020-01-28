package main

import (
	"fmt"
	"math"
)

// 関数も渡せる
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Goの関数はクロージャ
// それ自身の外部から変数を参照する関数
func adder() func(int) int { // 関数を返している
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func function_closures() {
	pos, neg := adder(), adder()
	fmt.Println("pos:", pos, "neg:", neg) // 関数のアドレスが表示されている
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	function_closures()
}

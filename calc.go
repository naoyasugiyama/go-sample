package main

import "fmt"

/*
 コメント文 C,C++と同じ
 セミコロンは自動補完
*/

func sample01() {
	// 変数宣言の仕方
	// 型名 変数名
	var num int
	var pow int
	// 宣言と同時に初期化で型を省略できる
	var result = 1

	num = 2
	pow = 4

	// この時は一行で二つ以上の文を書くときは必要
	var a = 4
	var b = 45

	for i := 0; i < pow; i++ {
		result *= num
	}

	fmt.Println("%dの%d乗は%dです.\n", num, pow, result)

	// if,for, switchの条件部にかっこは必要ない
	if a < 5 {
		fmt.Println("a < 5 : a=%d\n", a)
	}

	var c = 0 //  c declared and not used 使わないとエラー?

	// ifブロック内が短文でも波括弧は必須
	if b < 5 {
		fmt.Println(b)
		fmt.Println(c + b)
	}

}

func sample02() {
	//type
	var flag bool
	flag = false

	if flag {
		fmt.Print("flag check\n")
	}

	// 型
	/*
		var unsignedInt8 uint8
		unsignedInt8 = 255 // max
		var unsignedInt16 uint16
		unsignedInt16 = 65535
		var unsignedInt32 uint32
		unsignedInt32 = 4294967295
		var unsignedInt64 uint64
		unsignedInt64 = 18446744073709551615
		var unsingedInt uint // 環境依存
		var uint_ptr uintptr // 環境依存 ptr

		// 符号付き
		int8 -128 ~ 128
		int16 -32768 ~ 32767
		int32 -2147483648 ～ 2147483647
		int64 -9223372036854775808 ～ 9223372036854775807
		int int32 or int64

		//float
		float32 IEEE-754 32bit
		float64 IEEE-754 64bit

		//複数数
		complex64 　　?
		complex128

		//その他
		byte uint8のエイリアス
		rune int32のエイリアス

		//文字列型
		string

		一文字目が小文字の場合は、そのパッケージだけのスコープ変数
		一文字目が大文字の場合は、他のパッケージでも見えるスコープ変数
	*/

	// 算術演算
	// 二項演算子
	// ほとんどC++とかcと同じ
	// + - * /
	// % 余り
	// & 論理積
	// | ビット論理和
	// ^ ビット排他的論理和
	// &^ 各ビットの論理積の否定
	// << 左辺の算術左シフト
	// >> 左辺の算術右シフト

	var pos = 10
	pos += 2
	fmt.Println(pos)
	pos <<= 3
	fmt.Println(pos)

	fmt.Println(^pos)

	pos++
	fmt.Println(pos)
	// 比較演算子、論理演算子、アドレス演算子 C++,Cと同じ
	// アドレス演算子
	fmt.Println(&pos)
	//	fmt.Println(*pos)

	// 代入・特殊 左辺の変数が定義済みの場合は使用できない
	test := 100
	fmt.Println(test)
}

// 使い方,
func sample03() {
	type Vertex struct {
		x float32
		y float32
		z float32
	}

	v := Vertex{0.0, 0.0, 0.0}
	fmt.Println(v.x, v.y, v.z)

	var test [2]string
	test[0] = "Hello"
	test[1] = "World"

	fmt.Println(test[0] + test[1])

	var c complex128 = complex(1, 2) // 複素数に変換
	var fr float64 = real(c)
	var fi float64 = imag(c)

	fmt.Println(fr, fi) // 表示[1 2]
}

//送受信演算子
func sample04() {
	fmt.Println("送受信演算子説明")

	/*
		chan T
		chan<- float64
		<-chan int

		chan<- chan int
		chan<- <-chan int
		<-chan <-chan int
		chan(<-chan int)

		make(chan int, 100)
	*/
}

func main() {

	//	sample01()
	sample02()
	sample03()
}

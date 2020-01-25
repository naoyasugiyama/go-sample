package main

import "fmt"

func sample() {
	var street_address string
	street_address = "data center"
	fmt.Println("oosaka", street_address)

	// 宣言と同時に初期化で型名省略
	var user_name = "nsugiyama"
	fmt.Println("Hello,", user_name)
	// 関数内に限り、 := 演算子で代入するとvarキーワードも省略可能
	os_name := "windows"
	fmt.Println("user os,", os_name)
}

// 定数
func sample01() {
	const title = "Go言語入門"
	fmt.Println(title)
}

// リテラル
func sample02() {
	var decimal_number int = 1234 // 10
	var octal_number int = 073    // 8
	var hexadecima int = 0xA3     // 16

	fmt.Printf("dacial: [%d]\n", decimal_number)
	fmt.Printf("octal: [%o]\n", octal_number)
	fmt.Printf("hexadecima: [%x]\n", hexadecima)

	// 浮動小数リテラル
	// 10進数のみ
	var pI = 3.1415
	var val1 = .25     // 0.25
	var val2 = 12.     // 12.
	var val3 = 1.25e-3 // 指数表記
	fmt.Printf("PI: [%f]\n", pI)
	fmt.Printf("val1: [%f]\n", val1)
	fmt.Printf("val2: [%f]\n", val2)
	fmt.Printf("val3: [%f]\n", val3)

	// 虚数リテラル
	var complex01 = 2i
	var complex02 = 3.14151i
	var complex03 = 1.25e-3i

	fmt.Println("complex01:", complex01)
	fmt.Println("complex02:", complex02)
	fmt.Println("complex03:", complex03)

	// 複数型と虚数とは？
	// https://ja.wikipedia.org/wiki/%E8%A4%87%E7%B4%A0%E6%95%B0%E5%9E%8B

	// ルーンリテラル
	// Unicodeのコードポイントを表現する整数のことを(Rune)と呼ぶ
	// ルーンリテラルはルーン１つを表現するリテラル
	// \u
	var rune_a = 'a'
	var rune_a01 = 'あ'
	var rune_nl = '\n'
	var rune_point = '\u12AB'

	fmt.Printf("[%s],[%s],[%s]", rune_a, rune_a01, rune_nl)
	fmt.Printf("%s\n", rune_point)

	/*
	 二種類の表現方法がある
	 raw文字列リテラル
	 `` バッククオートで囲まれた文字列をraw文字数リテラルと呼ぶ
	 raw文字数リテラル内ではいかなるエスケープシーケンス・コードポイントも評価されない
	 ただし、リテラル内でバッククオートを利用することはできない
	*/
	var rune_abc = `abc`
	var rune_aa_nl = `\n` // 改行ではなく、\とnの二文字として扱う
	// 前の行と合わせて、改行を含む１つの文字列として扱う
	var rune_ab = `
	------------------------- 
		test code 
	------------------------- 
	`

	fmt.Printf("[%s]\n[%s]\n[%s]\n", rune_abc, rune_aa_nl, rune_ab)

	/*
		interpreted 文字列リテラル
		ダブルクオートで囲まれた文字列はinterpreted文字列リテラルという
		interpreted文字列リテラル内ではエスケープシーケンス、
		Unicodeコードポイントが評価される
	*/

	var interpreted_abc = "abc"
	var interpreted_pos = "ab\ncd"
	var interpreted_aiu = "\u3042\u3044\u3046"
	var interpreted_xyz = "\u0078\u0079\u007a"

	fmt.Printf("[%s]\n", interpreted_abc)
	fmt.Printf("[%s]\n", interpreted_pos)
	fmt.Printf("[%s]\n", interpreted_aiu)
	fmt.Printf("[%s]\n", interpreted_xyz)

}

func main() {
	sample()
	sample01()
	sample02()
}

package main

import "fmt"

// typeで型の別名宣言ができる
type Score int

// 構造体
type Dictionary struct {
	name    string
	meaning string
}

// 仮想関数 定義
type ReadFunc func(Dictionary) string

func sample() {
	var readFunc ReadFunc
	var dict Dictionary
	readFunc = readOut
	dict.name = "ココア"
	dict.meaning = "カカオから"
	fmt.Println(readFunc(dict))
}

func readOut(dict Dictionary) string {
	return fmt.Sprintf("[%s]は[%s]生まれるッ!", dict.name, dict.meaning)
}

// 関数レシーバー
type quantity int

func (pty quantity) Show() { fmt.Printf("胡椒%dグラムです\n", pty) } // メソッド追加できる

func sample02() {
	var pepper quantity = 100
	pepper.Show()
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
	var myScore Score = 100
	fmt.Printf("点数は%d点です！\n", myScore)
	sample()
	sample02()
	sample03()
}

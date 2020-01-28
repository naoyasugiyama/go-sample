// map
package main

import (
	"fmt"
	"strings"
	//	"golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

/*
m = make(map[string]Vertex)
m["Bell Labs"] = Vertex{
	40.68433, -74.39967,
}
*/
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

//!> mapに渡すトップレベル型が単純な型名である場合は、リテラルの要素から推定できるので型名を省略することができる
//!> ↓の場合は Vertex { 13213, 213213 }の型名を省略している
var mk = map[string]Vertex{
	"Bell":   {40.6890, -74.399},
	"Google": {37.5456, -25.5681526},
}

//!> map要素に対する挿入や更新など
func mutating_map() {
	map_sp := make(map[string]int)
	// 挿入や更新
	map_sp["Answer"] = 42
	fmt.Println("The value:", map_sp["Answer"])
	//要素の取得
	answer := map_sp["Answer"]
	fmt.Println("The GetValue:", answer)
	//要素の削除
	delete(map_sp, "Answer")
	fmt.Println("The Value:", map_sp["Answer"])
	//要素の各五人
	v, ok := map_sp["Answer"] // 第二引数でbool値が入ってくる..
	fmt.Println("The Value:", v, "Present?", ok)
	if ok {
		fmt.Println("vを参照する処理...:", v) // みたいな形で使う
	}
}

func wordCounts(s string) map[string]int {

	reswords := make(map[string]int) // word

	words := strings.Fields(s)

	for _, word := range words {
		reswords[word] += 1
	}

	fmt.Println("reswords:", reswords)
	return reswords
}

func main() {
	fmt.Println(m)
	mutating_map()

	wordCounts("It has been raining since morning today.")
}

package main

import (
	"fmt"
	"runtime"
	"time"
)

func sample01() {
	hour := time.Now().Hour()

	if hour >= 6 && hour < 12 {
		fmt.Println("朝です")
	} else if hour < 19 {
		fmt.Println("昼です")
	} else {
		fmt.Println("夜です")
	}
}

// blackが省略される
func sample02() {
	dayOfWeek := "月"
	switch dayOfWeek {
	case "土":
		fmt.Println("休みです")
	case "日":
		fmt.Println("休みです")
	default:
		fmt.Println("仕事")
	}

	score := 65
	switch score {
	case 100:
		fmt.Println("All")
	case 70: // 明示的に fallthroughを呼べば blackはされない
		fallthrough
	case 50:
		fmt.Println("復習")
	default:
		fmt.Println("再勉強")
	}

	//!> カンマを使えば複数の条件を設定できる
	switch dayOfWeek {
	case "土", "日":
		fmt.Println("大体休み")
	default:
		fmt.Println("大体仕事")
	}

	//!> 条件分岐を書いてもできる
	hour := time.Now().Hour()
	switch {
	case hour >= 6 && hour < 12:
		fmt.Println("朝です")
	case hour < 19:
		fmt.Println("昼です")
	default:
		fmt.Println("夜です")
	}

}

// go for 文
func sample03() {
	print("get odd:[%d]", getoddCount1(100))

	// 無限ループ
	i := 0
	for {
		i++
		if i >= 100 {
			break
		} else if i/2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// コレクション内のループ
	dayOfWeeks := [...]string{"mon", "tue", "wed", "thu", "fri", "sta", "sun"}

	for idx, day := range dayOfWeeks {
		fmt.Printf("%d番目の曜日は%s曜日です。\n", idx+1, day)
	}
}

// 奇数カウント
func getoddCount1(number int) int {
	count := 0
	for i := 0; i <= number; i++ {
		if i%2 != 0 {
			count++
		}
	}
	return count
}

// get os
func getOs() {
	os := runtime.GOOS
	fmt.Println(os)
}

//
func main() {
	//	sample01()
	//	sample02()
	sample03()
	getOs()
}

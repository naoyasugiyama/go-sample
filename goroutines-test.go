package main

import (
	"fmt"
	"time"
)

/*
 Goroutines
 Goのランタイムに管理される軽量なスレッド
 go function(x, y) で新しいgoroutineが実行される
 function x,y の評価は実行元(current)のgoroutineで実行され、functionの実行は
 新しいgoroutineで実行される

 goroutineは同じアドレス空間で実行される
 共有メモリへのアクセスは必ず同期する必要がある Channelsが必要

*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("10")
	say("11")
}

package main

import (
	"fmt"
	"time"
)

/*
	select
	goroutineを複数の通信操作で待たせることができる
	複数あるcaseのいずれかが準備できるようになるまでブロックする
	準備ができたcaseを実行する
	複数のcaseの準備ができている場合は、ランダムに選択される

	どのcaseも準備ができてないとき、selectの中からdefaultが実行される.


*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func defualtSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

	// default selection
	defualtSelection()
}

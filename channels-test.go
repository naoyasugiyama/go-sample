package main

import "fmt"

/*
チャンネルオペレータの<-を用いて値の送受信ができる

ch <- v //vをchannel　chへ送信する
v := <-ch // chから受診した変数をvへ割り当てる

// 生成方法
ch := make(chan int)

通常、片方が準備できるまで送受信はブロックされる。これにより、明確なロックや条件変更がなくても
goroutineの同期を可能にする


*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

/*
 チャンネルはバッファ(buffer)として利用できる
 バッファを持つチャンネルを初期化するには、makeの二つ目の引数にバッファの長さを入れる
 ※バッファが詰まった時は、チャンネルへの送信をブロックする。バッファが空の時は、チャンネルの送信をブロックする
 バッファが詰まるようにサンプルコードを変更し、何が起こるか見てみる
*/

func bufferchannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.bufferchannels()
			あってるのか？
	*/
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
チャネルは、ファイルとは異なり、通常は、closeする必要はない 
closeするのは、これ以上値が来ないことを受け手が知る必要があるとき。
 例えば、 range ループを終了するという場合です。
*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)     // 作ったチャンネルを閉じることができる
	v, ok := <-c // これで確認することができる
	fmt.Printf("(%v, %v)\n", v, ok
}

func range_and_close() {
	c0 := make(chan int, 10)
	go fibonacci(cap(c0), c0)
	for i := range c0 {
		fmt.Println(i)
	}
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //receive from c

	fmt.Println(x, y, x+y)

	bufferchannels()
	fmt.Println("----------------------")
	range_and_close()
}

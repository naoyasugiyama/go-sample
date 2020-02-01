package main

import "fmt"

/*
Stringer
 もっともよくつかわれる interface の一つ
 %GOPATH%\Go\src\strings
 type Stringer interface {
	 String() string
 }
 変数を文字列で出力するために使う
*/

type Product struct {
	Name string
	Age  int
}

func (p Product) String() string {
	return fmt.Sprintf("%v (%v yers)", p.Name, p.Age)
}

// ip Adder
type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

func main() {
	a := Product{"Asohe 〇uper Day", 192}
	b := Product{"olio e peperoncino", 980}
	fmt.Println(a, b)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

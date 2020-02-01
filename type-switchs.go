package main

// sample dirに入っているから自動変換でsampleに変換される vscodeのgo autoコンプリートの設定か

import "fmt"

/*
 空のインターフェイスには.(type)で型タイプを見ることが出来る
*/
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Int Type Twice: %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	case float64:
		fmt.Printf("Float %v is %T\n", v, v)
	default:
		fmt.Printf("I don't konw about type %T !\n", v)
	}
}

func main() {
	do(3.1415)
	do("Waoooooo!")
	do(true)
	do(65535)
	do(nil)
	do("") // zero string if "string-type" == "" が可能だと...

}

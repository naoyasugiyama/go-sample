package main

import (
	"fmt"
)

/*
ラベル　仕様
*/

func sample07() {
FOR_LABEL:
	for i := 0; i < 10; i++ {
		switch {
		case i == 5:
			break FOR_LABEL
		default:
			fmt.Println(i)
		}
	}
}

func sample08() {
LABEL1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 0 && j == 5 {
				// 1番目のforへcontinue
				//				fmt.Println("LABEL1")
				continue LABEL1
			} else if i == 1 && j == 1 {
				continue
			}
			//			fmt.Printf("i:[%d] j:[%d]\n", i, j)
		}
	}

	// goto LABEL2
	for k := 0; k < 10; k++ {
		fmt.Println(k)
		if k == 2 {
			goto LABEL2
		}
	}

LABEL2:
	fmt.Println("Sample08 End")
}

func main() {
	fmt.Println("label")
	//	sample07()
	sample08()
}

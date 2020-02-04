package main

import "fmt"

/*
/import "golang.org/x/tour/tree"
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree
func (t *Tree) String() string
*/


/*
 Exercise: Equivalent Binary Trees

  Walk walks the tree t sending all values
  from the tree to the channel ch.
*/
func Walk(t *Tree, ch chan int) {
	var walker func(t *Tree)

	walker = func(t *Tree) {
		if t == nil {
			return 
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	colse(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan []int,2)

	go walk(t1,ch1[0])
	go walk(t2,ch1[1])

	for value:= range ch1 {
		if value != <-ch1[1] {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i ++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

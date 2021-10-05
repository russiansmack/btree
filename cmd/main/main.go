package main

import "fmt"
import "github.com/russiansmack/btree"

func main() {
	b := btree.NewBTree()

	test := []int{12,11,90,82,7,9}
	b.InsertMany(test)

	Stringify(b.Top, 0)
	found, depth := b.FindDeepest()
	fmt.Println("deepest", found.Value, "depth", depth)
}

func Stringify(n *btree.Node, depth int) {
	if n != nil {
		format := ""
		for i := 0; i < depth; i++ {
			format += "       "
		}
		format += "----["
		depth++
		Stringify(n.Left, depth)
		fmt.Printf(format+"%d\n", n.Value)
		Stringify(n.Right, depth)
	}
}


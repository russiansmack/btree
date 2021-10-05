package btree

type BTree struct {
	Top *Node
}

func NewBTree() *BTree {
	return &BTree{Top: nil}
}

func (b *BTree) Insert(v int) {
	if b.Top == nil {
		b.Top = NewNode(v)
		return
	}

	b.Top.Insert(v)
}

func (b *BTree) InsertMany(values []int) {
	for _, v := range values {
		b.Insert(v)
	}
}

func (b *BTree) FindDeepest() (*Node, int) {
	return b.Top.FindDeepest(0)
}

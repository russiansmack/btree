package btree

type Node struct {
	Value int
	// Always smaller than Value
	Left *Node
	// Always bigger than Value
	Right *Node
}

func NewNode(v int) *Node {
	return &Node{Value: v}
}

func (n *Node) Insert(v int) {
	// Do nothing if we have a duplicate node
	if v == n.Value {
		return
	}

	if v < n.Value {
		if n.Left == nil {
			n.Left = NewNode(v)
		} else {
			n.Left.Insert(v)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(v)
		} else {
			n.Right.Insert(v)
		}
	}
}

func (n *Node) FindDeepest(startingDepth int) (*Node, int) {
	var lNode, rNode *Node
	var lDepth, rDepth int
	var hasChild bool

	if n.Left != nil {
		hasChild = true
		lNode, lDepth = n.Left.FindDeepest(startingDepth + 1)
	}

	if n.Right != nil {
		hasChild = true
		rNode, rDepth = n.Right.FindDeepest(startingDepth + 1)
	}

	// Found our deepest node which happens to be leaf (without children)
	if !hasChild {
		return n, startingDepth
	}

	if lDepth > rDepth {
		return lNode, lDepth
	} else {
		// if they are equal or rDepth is more on the Right return Right node
		return rNode, rDepth
	}
}

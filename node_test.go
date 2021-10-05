package btree

import (
	"math"
	"math/rand"
	"testing"
)

func TestNewNode(t *testing.T) {
	if NewNode(1).Value != 1 {
		t.Fatal("NewNode fails to create nodes")
	}
}

func TestNode_Insert(t *testing.T) {
	root := NewNode(2)
	root.Insert(3)
	root.Insert(1)

	if root.Right.Value != 3 {
		t.Fatal("Insert Right didn't work")
	}
	if root.Left.Value != 1 {
		t.Fatal("Insert Left didn't work")
	}
}

func benchmarkNode_Insert(max int, b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		s = append(s, rand.Intn(max))
	}
	b.ResetTimer()

	root := NewNode(1)
	for n := 0; n < b.N; n++ {
		root.Insert(s[n])
	}
}

func BenchmarkNode_Insert1(b *testing.B)      { benchmarkNode_Insert(1, b) }
func BenchmarkNode_Insert5(b *testing.B)      { benchmarkNode_Insert(5, b) }
func BenchmarkNode_Insert10(b *testing.B)     { benchmarkNode_Insert(10, b) }
func BenchmarkNode_Insert100(b *testing.B)    { benchmarkNode_Insert(100, b) }
func BenchmarkNode_Insert100000(b *testing.B) { benchmarkNode_Insert(100000, b) }

func TestNode_FindDeepest(t *testing.T) {
	b := NewNode(8)
	b.Insert(4)
	b.Insert(10)
	b.Insert(2)
	b.Insert(6)
	b.Insert(1)
	b.Insert(3)
	b.Insert(5)
	b.Insert(7)
	b.Insert(9)
	b.Insert(11)

	found, depth := b.FindDeepest(0)

	if found.Value != 7 || depth != 3 {
		t.Fatalf(`FindDeepest(0) = depth:%v, Value:%v, want "depth:%v, Value:%v", error`, depth, found.Value, 3, 7)
	}
}

func TestNode_FindDeepestVeryDeep(t *testing.T) {
	root := NewNode(1)
	for n := 0; n < math.MaxInt32/200; n++ {
		root.Insert(rand.Intn(math.MaxInt32))
	}

	root.FindDeepest(0)
}

func BenchmarkNode_FindDeepestSimple(b *testing.B) {
	n := NewNode(8)
	n.Insert(4)
	n.Insert(10)
	n.Insert(2)

	b.ResetTimer()
	n.FindDeepest(0)
}

func benchmarkNode_FindDeepest(max int, b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		s = append(s, rand.Intn(max))
	}

	root := NewNode(1)
	for n := 0; n < b.N; n++ {
		root.Insert(s[n])
	}
	b.ResetTimer()
	root.FindDeepest(0)
}

func BenchmarkNode_FindDeepest1(b *testing.B)      { benchmarkNode_FindDeepest(1, b) }
func BenchmarkNode_FindDeepest5(b *testing.B)      { benchmarkNode_FindDeepest(5, b) }
func BenchmarkNode_FindDeepest10(b *testing.B)     { benchmarkNode_FindDeepest(10, b) }
func BenchmarkNode_FindDeepest100(b *testing.B)    { benchmarkNode_FindDeepest(100, b) }
func BenchmarkNode_FindDeepest100000(b *testing.B) { benchmarkNode_FindDeepest(100000, b) }

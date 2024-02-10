package redblacktree

import (
	"fmt"
	"testing"
)

// IntComparator provides a basic comparison on int
func IntComparator(a, b interface{}) int {
	aAsserted := a.(int)
	bAsserted := b.(int)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func TestRedBlackTreePut(t *testing.T) {
	tree, err := NewWith(IntComparator)
	defer tree.db.CloseDB()
	if err != nil {
		t.Error("could not initialize")
	}
	var key [8]byte
	copy(key[:], "00000012")
	tree.Put(key, RedBlackTreeNodeDBValue{Value: "x", Key: 120}) // 1->x
	var key2 [8]byte
	copy(key2[:], "00000002")
	tree.Put(key2, RedBlackTreeNodeDBValue{Value: "b", Key: 112}) // 1->x, 2->b (in order)
	var key3 [8]byte
	copy(key3[:], "00000003")
	tree.Put(key3, RedBlackTreeNodeDBValue{Value: "ba", Key: 12})
	// tree.Put(1, RedBlackTreeNodeDBValue{Value: "a"}) // 1->a, 2->b (in order, replacement)
	// tree.Put(3, RedBlackTreeNodeDBValue{Value: "c"}) // 1->a, 2->b, 3->c (in order)
	// tree.Put(4, RedBlackTreeNodeDBValue{Value: "d"}) // 1->a, 2->b, 3->c, 4->d (in order)
	// tree.Put(5, RedBlackTreeNodeDBValue{Value: "e"}) // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	// tree.Put(6, RedBlackTreeNodeDBValue{Value: "f"}) // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)
	t.Log(tree.Size, "sep")
	fmt.Println(tree)
	//
	//  RedBlackTree
	//  │           ┌── 6
	//  │       ┌── 5
	//  │   ┌── 4
	//  │   │   └── 3
	//  └── 2
	//
}

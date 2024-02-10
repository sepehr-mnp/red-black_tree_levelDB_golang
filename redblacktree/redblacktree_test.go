package redblacktree

import (
	"fmt"
	"testing"
)

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
	copy(key[:], "1")
	tree.Put(key, RedBlackTreeNodeDBValue{Value: "x"}) // 1->x
	copy(key[:], "1")
	tree.Put(key, RedBlackTreeNodeDBValue{Value: "b"}) // 1->x, 2->b (in order)
	// tree.Put(1, RedBlackTreeNodeDBValue{Value: "a"}) // 1->a, 2->b (in order, replacement)
	// tree.Put(3, RedBlackTreeNodeDBValue{Value: "c"}) // 1->a, 2->b, 3->c (in order)
	// tree.Put(4, RedBlackTreeNodeDBValue{Value: "d"}) // 1->a, 2->b, 3->c, 4->d (in order)
	// tree.Put(5, RedBlackTreeNodeDBValue{Value: "e"}) // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	// tree.Put(6, RedBlackTreeNodeDBValue{Value: "f"}) // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)
	t.Log(tree)
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

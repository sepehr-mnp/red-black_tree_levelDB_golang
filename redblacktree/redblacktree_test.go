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
	tree.Load()
	// defer tree.Save()
	// var key [8]byte
	// copy(key[:], "00000012")
	// tree.Put(key, RedBlackTreeNodeDBValue{Value: "x", Key: 120}) // 1->x
	// copy(key[:], "00000002")
	// tree.Put(key, RedBlackTreeNodeDBValue{Value: "b", Key: 112}) // 1->x, 2->b (in order)

	// copy(key[:], "00000003")
	// tree.Put(key, RedBlackTreeNodeDBValue{Value: "ba", Key: 12})
	// copy(key[:], "00000004")
	// tree.Put(key, RedBlackTreeNodeDBValue{Value: "bajaja", Key: 1})
	// copy(key[:], "00000005")
	// tree.Put(key, RedBlackTreeNodeDBValue{Value: "bagaa", Key: 166})
	// copy(key[:], "00000006")
	// tree.Put(key, RedBlackTreeNodeDBValue{Value: "poapba", Key: 1212})
	// copy(key[:], "00000007")
	// tree.Put(key, RedBlackTreeNodeDBValue{Value: "lallaba", Key: 18})
	t.Log(tree.Size)
	fmt.Println(tree)

}

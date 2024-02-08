package redblacktree

import (
	"testing"
)

// / run: go test -v .
func TestRedBlackTreeValueEncode(t *testing.T) {
	t.Log(&RedBlackTreeNode{})
}

func TestNilByteArray(t *testing.T) {
	if (RedBlackTreeNode{}).DBKey != nilByteArray {
		t.Error("were not match", RedBlackTreeNode{}.DBKey, nilByteArray)
	}
	t.Log("was OK :)(i hate this fake smile but i needed sth to make difference between test package restults and mine but while writing this i relized there were infinity options that were not as stupid as this but i wont change this so maby it brings you laughter in the future and also because i seek attention.)")
}

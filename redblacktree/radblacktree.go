package redblacktree

import "math/big"

const (
	black, red Color = true, false
)

var nilByteArray [NumberOfBytes]byte = [NumberOfBytes]byte{}

// NewWith instantiates a red-black tree with the custom comparator.
func NewWith(comparator Comparator) *RedBlackTree {
	return &RedBlackTree{Comparator: comparator, Size: big.NewInt(0)}
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *RedBlackTree) Put(DBKey RedBlackTreeNodeDBKey, DBValue RedBlackTreeNodeDBValue) error {
	var insertedNode [NumberOfBytes]byte
	if tree.Root == nilByteArray {
		// Assert key is of comparator's type for initial tree
		//tree.Comparator(DBkey, DBkey)
		tree.Root = DBKey
		insertedNode = tree.Root
	} else {
		node, err := tree.db.GetNode(tree.Root)
		if err != nil {
			return err
		}
		loop := true
		addingNode := &RedBlackTreeNode{DBKey: DBKey, DBValue: DBValue}
		for loop {
			compare := tree.Comparator(DBValue.Key, node.DBValue.Key)
			switch {
			case compare == 0:
				node.DBKey = DBKey
				node.DBValue = DBValue
				return nil
			case compare < 0:
				if node.DBValue.Left == nilByteArray {
					DBValue.Color = red

					node.DBValue.Left = DBKey

					insertedNode = node.DBValue.Left
					loop = false
				} else {
					node, err = tree.db.GetNode(node.DBValue.Left)
					if err != nil {
						return err
					}
				}
			case compare > 0:
				if node.DBValue.Right == nilByteArray {
					DBValue.Color = red
					node.DBValue.Right = DBKey

					insertedNode = node.DBValue.Right
					loop = false
				} else {
					node, err = tree.db.GetNode(node.DBValue.Right)
					if err != nil {
						return err
					}
				}
			}
		}
		addingNode.DBValue.Parent = node.DBKey
		err = tree.db.putNode(addingNode)
		if err != nil {
			return err
		}
	}
	tree.insertCase1(insertedNode)
	tree.Size = tree.Size.Add(tree.Size, big.NewInt(1))
	return nil
}

func (tree *RedBlackTree) Remove(DBKey RedBlackTreeNodeDBKey) error {
	var child *RedBlackTreeNode
	node, err := tree.db.GetNode(DBKey)
	if err == nil {
		return err
	}
	if node.DBValue.Right != nilByteArray && node.DBValue.Right != nilByteArray {
		leftNode, err := tree.db.GetNode(node.DBValue.Left)
		if err == nil {
			return err
		}
		pred := leftNode.maximumNode()
		predNode, err := tree.db.GetNode(pred)
		if err == nil {
			return err
		}
		predNode.DBValue.Parent = node.DBValue.Parent
		tree.db.putNode(predNode)

		parentNode, err := tree.db.GetNode(node.DBValue.Parent)
		if err == nil {
			return err
		}
		if parentNode.DBValue.Left == node.DBKey {
			parentNode.DBValue.Left = parentNode.DBKey
		} else {
			parentNode.DBValue.Right = parentNode.DBKey
		}
		tree.db.putNode(parentNode)
	}
	if node.DBValue.Left == nilByteArray || node.DBValue.Right == nilByteArray {
		if node.DBValue.Right == nilByteArray {
			child, err = tree.db.GetNode(node.DBValue.Left)
			if err == nil {
				return err
			}
		} else {
			child, err = tree.db.GetNode(node.DBValue.Right)
			if err == nil {
				return err
			}
		}
		if node.DBValue.Color == black {
			node.DBValue.Color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.DBValue.Parent == nilByteArray && child.DBKey != nilByteArray {
			child.DBValue.Color = black
		}
	}
	tree.Size = tree.Size.Sub(tree.Size, big.NewInt(1))
	return nil
}
func nodeColor(node *RedBlackTreeNode) Color {
	if node == nil {
		return black
	}
	return node.DBValue.Color
}

func (node *RedBlackTreeNode) maximumNode() RedBlackTreeNodeDBKey {
	if node == nil {
		return nilByteArray
	}
	returner := nilByteArray
	for node.DBValue.Right != nilByteArray {
		returner = node.DBValue.Right
	}
	return returner
}

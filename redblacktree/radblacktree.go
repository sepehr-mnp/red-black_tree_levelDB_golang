package redblacktree

import (
	"errors"
	"fmt"
	"math/big"
)

const (
	black, red Color = true, false
)

var nilByteArray [NumberOfBytes]byte = [NumberOfBytes]byte{}

// NewWith instantiates a red-black tree with the custom comparator.
func NewWith(comparator Comparator) (*RedBlackTree, error) {
	dbGotten, err := GetNewLvelDBDatabase("db/")
	if err != nil {
		return nil, err
	}
	return &RedBlackTree{Comparator: comparator, Size: big.NewInt(0), db: dbGotten}, nil
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
		tree.db.putNode(&RedBlackTreeNode{DBKey: DBKey, DBValue: DBValue})
		fmt.Println(1, " sep")
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
				err = tree.db.putNode(node)
				if err != nil {
					return err
				}
				return nil
			case compare < 0:
				if node.DBValue.Left == nilByteArray {
					DBValue.Color = red

					node.DBValue.Left = DBKey

					insertedNode = node.DBValue.Left

					loop = false
					err = tree.db.putNode(node)
					if err != nil {
						return err
					}
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
					err = tree.db.putNode(node)
					if err != nil {
						return err
					}
				} else {
					node, err = tree.db.GetNode(node.DBValue.Right)
					if err != nil {
						return err
					}
				}
			}
		}
		addingNode.DBValue.Parent = node.DBKey
		fmt.Println(addingNode)
		err = tree.db.putNode(addingNode)
		if err != nil {
			return err
		}
		fmt.Println(2, " sep")
	}
	insertedNodeGottenDB, _ := tree.db.GetNode(insertedNode)
	tree.insertCase1(insertedNodeGottenDB)
	tree.Size = tree.Size.Add(tree.Size, big.NewInt(1))
	return nil
}

// func (tree *RedBlackTree) Remove(DBKey RedBlackTreeNodeDBKey) error {
// 	var child *RedBlackTreeNode
// 	node, err := tree.db.GetNode(DBKey)
// 	if err == nil {
// 		return err
// 	}
// 	if node.DBValue.Right != nilByteArray && node.DBValue.Right != nilByteArray {
// 		leftNode, err := tree.db.GetNode(node.DBValue.Left)
// 		if err != nil {
// 			return err
// 		}
// 		pred := leftNode.maximumNode()
// 		predNode, err := tree.db.GetNode(pred)
// 		if err != nil {
// 			return err
// 		}
// 		predNode.DBValue.Parent = node.DBValue.Parent
// 		tree.db.putNode(predNode)

//			parentNode, err := tree.db.GetNode(node.DBValue.Parent)
//			if err != nil {
//				return err
//			}
//			if parentNode.DBValue.Left == node.DBKey {
//				parentNode.DBValue.Left = parentNode.DBKey
//			} else {
//				parentNode.DBValue.Right = parentNode.DBKey
//			}
//			tree.db.putNode(parentNode)
//		}
//		if node.DBValue.Left == nilByteArray || node.DBValue.Right == nilByteArray {
//			if node.DBValue.Right == nilByteArray {
//				child, err = tree.db.GetNode(node.DBValue.Left)
//				if err != nil {
//					return err
//				}
//			} else {
//				child, err = tree.db.GetNode(node.DBValue.Right)
//				if err != nil {
//					return err
//				}
//			}
//			if node.DBValue.Color == black {
//				node.DBValue.Color = nodeColor(child)
//				tree.deleteCase1(node)
//			}
//			tree.replaceNode(node, child)
//			if node.DBValue.Parent == nilByteArray && child.DBKey != nilByteArray {
//				child.DBValue.Color = black
//			}
//		}
//		tree.Size = tree.Size.Sub(tree.Size, big.NewInt(1))
//		return nil
//	}
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

// Left returns the left-most (min) node or nil if tree is empty.
func (tree *RedBlackTree) Left() (*RedBlackTreeNode, error) {
	var parent *RedBlackTreeNode
	current, err := tree.db.GetNode(tree.Root)
	if err != nil {
		return nil, err
	}
	for current != nil {
		parent = current
		current, err = tree.db.GetNode(current.DBValue.Left)
		if err != nil {
			return nil, err
		}
	}
	return parent, nil
}

func (tree *RedBlackTree) lookup(key interface{}) (*RedBlackTreeNode, error) {
	node := tree.Root
	for node != nilByteArray {
		currentNode, err := tree.db.GetNode(node)
		if err != nil {
			return nil, err
		}
		compare := tree.Comparator(key, currentNode.DBKey)
		switch {
		case compare == 0:
			return currentNode, nil
		case compare < 0:
			node = currentNode.DBValue.Left
		case compare > 0:
			node = currentNode.DBValue.Right
		}
	}
	return nil, errors.New("not found")
}

func (tree *RedBlackTree) grandparent(node *RedBlackTreeNode) *RedBlackTreeNode {
	if node != nil && node.DBValue.Parent != nilByteArray {
		currentNodeParent, _ := tree.db.GetNode(node.DBValue.Parent)
		currentNodeGrandParent, _ := tree.db.GetNode(currentNodeParent.DBValue.Parent)
		return currentNodeGrandParent

	}
	return nil
}

func (tree *RedBlackTree) uncle(node *RedBlackTreeNode) *RedBlackTreeNode {
	if node == nil || node.DBValue.Parent == nilByteArray || tree.grandparent(node) == nil {
		return nil
	}
	parent, _ := tree.db.GetNode(node.DBValue.Parent)
	return tree.uncle(parent)
}

func (tree *RedBlackTree) sibling(node *RedBlackTreeNode) *RedBlackTreeNode {
	if node == nil || node.DBValue.Parent == nilByteArray {
		return nil
	}
	parrent, _ := tree.db.GetNode(node.DBValue.Parent)
	if node.DBKey == parrent.DBValue.Left {
		parrentRightChild, _ := tree.db.GetNode(parrent.DBValue.Right)
		return parrentRightChild
	}
	parrentLeftChild, _ := tree.db.GetNode(parrent.DBValue.Left)
	return parrentLeftChild
}

func (tree *RedBlackTree) rotateLeft(node *RedBlackTreeNode) {
	right, _ := tree.db.GetNode(node.DBValue.Right)
	tree.replaceNode(node, right)
	node.DBValue.Right = right.DBValue.Left
	if right.DBValue.Left != nilByteArray {
		leftOfRight, _ := tree.db.GetNode(right.DBValue.Left)
		leftOfRight.DBValue.Parent = node.DBKey
	}
	right.DBValue.Left = node.DBKey
	node.DBValue.Parent = right.DBKey
	tree.db.putNode(right)
	tree.db.putNode(node)
}

func (tree *RedBlackTree) rotateRight(node *RedBlackTreeNode) {
	left, _ := tree.db.GetNode(node.DBValue.Left)
	tree.replaceNode(node, left)
	node.DBValue.Left = left.DBValue.Right
	if left.DBValue.Right != nilByteArray {
		rightOfLeft, _ := tree.db.GetNode(left.DBValue.Right)
		rightOfLeft.DBValue.Parent = node.DBKey
	}
	left.DBValue.Right = node.DBKey
	node.DBValue.Parent = left.DBKey
	tree.db.putNode(left)
	tree.db.putNode(node)
}

func (tree *RedBlackTree) replaceNode(old *RedBlackTreeNode, new *RedBlackTreeNode) {
	if old.DBValue.Parent == nilByteArray {
		tree.Root = new.DBKey
	} else {
		parent, _ := tree.db.GetNode(old.DBValue.Parent)
		if old.DBKey == parent.DBValue.Left {
			parent.DBValue.Left = new.DBKey
		} else {
			parent.DBValue.Right = new.DBKey
		}
		tree.db.putNode(parent)
	}
	if new != nil {
		new.DBValue.Parent = old.DBValue.Parent
		tree.db.putNode(new)
	}
}

func (tree *RedBlackTree) insertCase1(node *RedBlackTreeNode) {
	if node.DBValue.Parent == nilByteArray {
		node.DBValue.Color = black
		tree.db.putNode(node)
	} else {
		tree.insertCase2(node)
	}
}

func (tree *RedBlackTree) insertCase2(node *RedBlackTreeNode) {
	parent, _ := tree.db.GetNode(node.DBValue.Right)
	if nodeColor(parent) == black {
		return
	}
	tree.insertCase3(node)
}

func (tree *RedBlackTree) insertCase3(node *RedBlackTreeNode) {
	uncle := tree.uncle(node)
	if nodeColor(uncle) == red {
		parent, _ := tree.db.GetNode(node.DBValue.Parent)
		parent.DBValue.Color = black
		uncle.DBValue.Color = black
		grandParent := tree.grandparent(node)
		grandParent.DBValue.Color = red
		tree.db.putNode(parent)
		tree.db.putNode(uncle)
		tree.db.putNode(grandParent)
		tree.insertCase1(grandParent)
	} else {
		tree.insertCase4(node)
	}
}

func (tree *RedBlackTree) insertCase4(node *RedBlackTreeNode) {
	grandParent := tree.grandparent(node)
	parent, _ := tree.db.GetNode(node.DBValue.Parent)
	if node.DBKey == parent.DBValue.Right && node.DBValue.Parent == grandParent.DBValue.Left {
		tree.rotateLeft(parent)
		node, _ = tree.db.GetNode(node.DBValue.Left) /// should i did the mappings too?
	} else if node.DBKey == parent.DBValue.Left && node.DBValue.Parent == grandParent.DBValue.Right {
		tree.rotateRight(parent)
		node, _ = tree.db.GetNode(node.DBValue.Right)
	}
	tree.insertCase5(node)
}

func (tree *RedBlackTree) insertCase5(node *RedBlackTreeNode) {
	parent, _ := tree.db.GetNode(node.DBValue.Parent)
	parent.DBValue.Color = black
	grandparent := tree.grandparent(node)
	grandparent.DBValue.Color = red
	if node.DBKey == parent.DBValue.Left && node.DBValue.Parent == grandparent.DBValue.Left {
		tree.rotateRight(grandparent)
	} else if node.DBKey == parent.DBValue.Right && node.DBValue.Parent == grandparent.DBValue.Right {
		tree.rotateLeft(grandparent)
	}
}

// / printing
// String returns a string representation of container
func (tree *RedBlackTree) String() string {
	str := "RedBlackTree\n"
	if tree.Size.Cmp(big.NewInt(0)) != 0 {
		treeRoot, _ := tree.db.GetNode(tree.Root)
		tree.output(treeRoot, "", true, &str)
	}
	return str
}

func (node *RedBlackTreeNode) String() string {
	return fmt.Sprintf("key: %v, right: %v, left: %v, parent, %v", node.DBKey, node.DBValue.Right, node.DBValue.Left, node.DBValue.Parent)
}

func (tree *RedBlackTree) output(node *RedBlackTreeNode, prefix string, isTail bool, str *string) {
	fmt.Println(3, "sep")
	if node.DBValue.Right != nilByteArray {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		rightNode, _ := tree.db.GetNode(node.DBValue.Right)
		tree.output(rightNode, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.DBValue.Left != nilByteArray {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		leftNode, _ := tree.db.GetNode(node.DBValue.Left)
		tree.output(leftNode, newPrefix, true, str)
	}
}

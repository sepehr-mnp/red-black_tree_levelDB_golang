package redblacktree

import "math/big"

const (
	NumberOfBytes int = 8
)

type Color bool

type Comparator func(a, b interface{}) int

type RedBlackTree struct {
	Root       [NumberOfBytes]byte
	db         *levelDB
	Comparator Comparator
	Size       *big.Int
}

type EncodableAndComparable interface {
	// /Comparator(a, b Encodable) int
	//Encodable
}

type Encodable interface {
	//Encode([]byte) error
	//Decode([]byte) error
}

type RedBlackTreeNode struct {
	DBKey   RedBlackTreeNodeDBKey
	DBValue RedBlackTreeNodeDBValue
}

type RedBlackTreeNodeDBValue struct {
	Parent [NumberOfBytes]byte
	Key    EncodableAndComparable
	Value  Encodable
	Color  Color
	Right  [NumberOfBytes]byte
	Left   [NumberOfBytes]byte
}

type RedBlackTreeNodeDBKey [NumberOfBytes]byte

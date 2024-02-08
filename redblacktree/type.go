package redblacktree

import "math/big"

const (
	NumberOfBytes int = 8
)

type Comparator[T any] func(x, y T) int

type RedBlackTree[K comparable] struct {
	Root       [NumberOfBytes]byte
	size       *big.Int
	db         *levelDB
	Comparator Comparator[K]
}
type Color bool

type Encodable interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

type RedBlackTreeNode struct {
	DBKey   RedBlackTreeNodeDBKey
	DBValue RedBlackTreeNodeDBValue
}

type RedBlackTreeNodeDBValue struct {
	Parent [NumberOfBytes]byte
	Key    Encodable
	Value  Encodable
	Color  Color
	Right  [NumberOfBytes]byte
	Left   [NumberOfBytes]byte
}

type RedBlackTreeNodeDBKey struct {
	DBKey [NumberOfBytes]byte
}

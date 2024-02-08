package redblacktree

const (
	NumberOfBytes int = 8
)

type RedBlackTree struct {
	root [NumberOfBytes]byte
	db   *levelDB
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

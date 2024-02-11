package redblacktree

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func (n *RedBlackTreeNode) Encode() ([]byte, error) {
	var data bytes.Buffer
	/// extract DBValue
	dbValue := &n.DBValue

	enc := gob.NewEncoder(&data)
	err := enc.Encode(*dbValue)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func (n *RedBlackTreeNode) Decode(data []byte) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	return dec.Decode(&n.DBValue)
}

func (n *RedBlackTreeNodeDBKey) Decode(data []byte) error {
	var key [NumberOfBytes]byte
	fmt.Println(data)
	copy(key[:], data)
	fmt.Println(88, key)
	//n = &RedBlackTreeNodeDBKey{key}
	return nil
}

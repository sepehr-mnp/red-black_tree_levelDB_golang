package redblacktree

import (
	"bytes"
	"encoding/gob"
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

func (t *RedBlackTree) Encode() ([]byte, error) {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(*t)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func (t *RedBlackTree) Decode(data []byte) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	return dec.Decode(&t)
}

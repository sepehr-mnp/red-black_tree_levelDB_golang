package redblacktree

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type levelDB struct {
	db *leveldb.DB
}

func GetNewLvelDBDatabase(path string) (*levelDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &levelDB{db}, nil
}

func (l *levelDB) CloseDB() {
	l.db.Close()
}

func (db *levelDB) putNode(node *RedBlackTreeNode) error {
	dbValue, err := node.Encode()
	if err != nil {
		return err
	}
	return db.db.Put(node.DBKey[:], dbValue, nil)
}

func (db *levelDB) GetNode(nodeKey RedBlackTreeNodeDBKey) (*RedBlackTreeNode, error) {
	bytesGotten, err := db.db.Get(nodeKey[:], nil)
	if err != nil {
		return nil, err
	}
	gottenNode := &RedBlackTreeNode{}
	err = gottenNode.Decode(bytesGotten)
	if err != nil {
		return nil, err
	}
	gottenNode.DBKey = nodeKey
	return gottenNode, nil
}

func (db *levelDB) DeleteNode(nodeKey RedBlackTreeNodeDBKey) error {
	return db.db.Delete(nodeKey[:], nil)
}

func (db *levelDB) Save(root RedBlackTreeNodeDBKey) error {
	fmt.Println(root)
	return db.db.Put([]byte("00000000"), root[:], nil) // reserved address for root
}

func (db *levelDB) Load() ([NumberOfBytes]byte, error) {
	byteArrayGotten, err := db.db.Get([]byte("00000000"), nil)
	fmt.Println(1, byteArrayGotten)
	//gottenDBKey := &RedBlackTreeNodeDBKey{}
	//err = gottenDBKey.Decode(byteArrayGotten)
	//fmt.Println(2, gottenDBKey)
	//return gottenDBKey, err
	var key [NumberOfBytes]byte
	copy(key[:], byteArrayGotten)
	return key, err
}

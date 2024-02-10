package redblacktree

import "github.com/syndtr/goleveldb/leveldb"

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

package index

import (
	"bitcask-go/data"
	"github.com/google/btree"
	"sync"
)

//使用google的btree
type BTree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

func NewBTree() *BTree {
	return &BTree{
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}
func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	item := &Item{
		key: key,
		pos: pos,
	}
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(item)
	bt.lock.Unlock()
	return true
}
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	item := &Item{key: key}
	bTreeItem := bt.tree.Get(item)
	if bTreeItem == nil {
		return nil
	}
	return bTreeItem.(*Item).pos
}
func (bt *BTree) Delete(key []byte) bool {
	item := &Item{key: key}
	bt.lock.Lock()
	bTreeItem := bt.tree.Delete(item)
	bt.lock.Unlock()
	if bTreeItem == nil {
		return false
	}
	return true
}

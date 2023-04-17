package index

import (
	"bitcask-go/data"
	"bytes"
	"github.com/google/btree"
)

//之后选用的数据结构实现这个接口即可
type Indexer interface {
	Put(key []byte, pos *data.LogRecordPos) bool
	Get(key []byte) *data.LogRecordPos
	Delete(key []byte) bool
}
type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (aI *Item) Less(bI btree.Item) bool {
	return bytes.Compare(aI.key, bI.(*Item).key) == -1
}

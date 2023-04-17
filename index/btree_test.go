package index

import (
	"bitcask-go/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBTree_Put(t *testing.T) {
	bTree := NewBTree()
	res := bTree.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 120})
	assert.True(t, res)

	res2 := bTree.Put(nil, &data.LogRecordPos{Fid: 50, Offset: 2})
	assert.True(t, res2)
}

func TestBTree_Get(t *testing.T) {
	bTree := NewBTree()
	res := bTree.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 120})
	assert.True(t, res)

	res2 := bTree.Put([]byte("a"), &data.LogRecordPos{Fid: 50, Offset: 2})
	assert.True(t, res2)
}

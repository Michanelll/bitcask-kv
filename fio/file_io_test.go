package fio

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestNewFileIOManager(t *testing.T) {
	fileIOManaer, err := NewFileIOManager(filepath.Join("D:\\Go\\src\\bitcask-go\\tmp", "a.data"))
	assert.Nil(t, err)
	assert.NotNil(t, fileIOManaer)
}

func TestFileIO_Write(t *testing.T) {
	fileIOManaer, err := NewFileIOManager(filepath.Join("D:\\Go\\src\\bitcask-go\\tmp", "a.data"))
	assert.Nil(t, err)
	assert.NotNil(t, fileIOManaer)

	n, err := fileIOManaer.Write([]byte("4531"))
	t.Log(n, err)
	n, err = fileIOManaer.Write([]byte("agrwre"))
	t.Log(n, err)
}

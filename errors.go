package bitcask_go

import "errors"

var (
	ErrKeyIsEmpty       = errors.New("key is empty")
	ErrIndexUpdatedFail = errors.New("update index fail")
	ErrKeyNotFound      = errors.New("key not found")
	ErrDataFileNotFound = errors.New("datafile not found")
)

package fio

const DataFilePerm = 0644

type IOManager interface {
	Read([]byte, int64) (int, error)
	Write([]byte) (int, error)
	//持久化数据
	Sync() error
	Close() error
}

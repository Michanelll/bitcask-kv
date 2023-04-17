package data

import "bitcask-go/fio"

type DataFile struct {
	FileId    uint32
	Offset    int64         //偏移量
	IOManager fio.IOManager //数据读写接口
}

func OpenFile(dirPath string, fileId uint32) (*DataFile, error) {
	return nil, nil
}
func (d *DataFile) Sync() error {
	return nil
}
func (d *DataFile) Write(buf []byte, offset int64) error {
	return nil
}
func (d *DataFile) ReadLogRecord(offset int64) (*LogRecord, error) {
	return nil, nil
}

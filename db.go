package bitcask_go

import (
	"bitcask-go/data"
	"bitcask-go/index"
	"sync"
)

type DB struct {
	option         Options                   //配置选项
	mu             *sync.RWMutex             //	并发控制
	activeDataFile *data.DataFile            //活跃文件
	oldDataFiles   map[uint32]*data.DataFile //不活跃文件
	index          index.Indexer
}

func (db *DB) Put(key []byte, value []byte) error {
	//首先判断key不为空
	if len(key) == 0 {
		return ErrKeyIsEmpty
	}
	//构造logRecord
	log_record := &data.LogRecord{
		Key:   key,
		Value: value,
		Type:  data.LogRecordNormal,
	}
	pos, err := db.appendLogRecord(log_record)
	if err != nil {
		return err
	}
	if ok := db.index.Put(key, pos); !ok {
		return ErrIndexUpdatedFail
	}
	return nil
}
func (db *DB) Get(key []byte) ([]byte, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	//首先判断key不为空
	if len(key) == 0 {
		return nil, ErrKeyIsEmpty
	}
	//取索引
	logRecordPos := db.index.Get(key)
	if logRecordPos == nil {
		return nil, ErrKeyNotFound
	}
	//去文件中找
	var dataFile *data.DataFile
	if db.activeDataFile.FileId == logRecordPos.Fid {
		dataFile = db.activeDataFile
	} else {
		dataFile = db.oldDataFiles[logRecordPos.Fid]
	}
	if dataFile == nil {
		return nil, ErrDataFileNotFound
	}
	logRecord, err := dataFile.ReadLogRecord(logRecordPos.Offset)
	if err != nil {
		return nil, err
	}
	//类型检查
	if logRecord.Type == data.LogRecordDeleted {
		return nil, ErrKeyNotFound
	}
	return logRecord.Value, nil
}

//追加写
func (db *DB) appendLogRecord(log_record *data.LogRecord) (*data.LogRecordPos, error) {
	//先加锁
	db.mu.Lock()
	defer db.mu.Unlock()
	//如果数据库刚完成初始化，就没有活跃文件
	if db.activeDataFile == nil {
		//设置新的活跃文件
		if err := db.setActiveDataFile(); err != nil {
			return nil, err
		}
	}
	//二进制编码
	encRecord, length := data.EncodeLogRecord(log_record)
	//如果超出活跃文件阈值，先持久化，再设置新的活跃文件
	if db.activeDataFile.Offset+length > db.option.ActiveFileThreshold {
		if err := db.activeDataFile.Sync(); err != nil {
			return nil, err
		}
	}
	db.oldDataFiles[db.activeDataFile.FileId] = db.activeDataFile
	if err := db.setActiveDataFile(); err != nil {
		return nil, err
	}
	offset := db.activeDataFile.Offset
	//写入
	if err := db.activeDataFile.Write(encRecord, offset); err != nil {
		return nil, err
	}
	//是否需要每次写后存
	if db.option.SyncsAfterWrites {
		if err := db.activeDataFile.Sync(); err != nil {
			return nil, err
		}
	}
	//构造索引
	pos := &data.LogRecordPos{
		Fid:    db.activeDataFile.FileId,
		Offset: offset,
	}
	return pos, nil
}
func (db *DB) setActiveDataFile() error {
	var intialFileId uint32 = 0
	if db.activeDataFile != nil {
		intialFileId = db.activeDataFile.FileId + 1
	}
	file, err := data.OpenFile(db.option.DirPath, intialFileId)
	if err != nil {
		return err
	}
	db.activeDataFile = file
	return nil
}

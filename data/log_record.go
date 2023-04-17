package data

type LogRecordType = byte

const (
	LogRecordNormal  = 0
	LogRecordDeleted = 1
)

//写入的记录
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

//内存索引，描述数据在磁盘的位置
type LogRecordPos struct {
	Fid    uint32 //文件id
	Offset int64
}

//返回byte和长度
func EncodeLogRecord(record *LogRecord) ([]byte, int64) {
	return nil, 0
}

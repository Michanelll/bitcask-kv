package bitcask_go

type Options struct {
	DirPath             string //数据目录
	ActiveFileThreshold int64  //活跃文件阈值
	SyncsAfterWrites    bool
}

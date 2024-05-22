package data

type LogRecordPos struct {
	Fid    uint32 // 文件 id 表示数据存储的文件位置
	Offset int64  // 偏移 表示将数据存储到的数据文件中的位置
}

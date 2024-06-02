package data

import "encoding/binary"

type LogRecordType byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
)

// crc type keySize valueSize
// 4 + 1 + 5 + 5 = 15
const maxLogRecordHeaderSize = binary.MaxVarintLen32*2 + 5

// LogRecord 写入到数据文件的记录
// 数据文件中的数据是追加写入的，类似日志的格式
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

// logRecordHeader LogRecord的头部信息
type logRecordHeader struct {
	crc        uint32        // crc 校验值
	recordType LogRecordType // 标识 LogRecord 的类型
	keySize    uint32        // key 的长度
	valueSize  uint32        // value 的长度
}

// LogRecordPos 数据内存索引，主要是描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件 id 表示数据存储的文件位置
	Offset int64  // 偏移 表示将数据存储到的数据文件中的位置
}

func EncodeLogRecord(logRecord *LogRecord) ([]byte, int64) {
	return nil, 0
}

// 对字节数组中的 Header 信息进行解码
func decodeLogRecordHeader(buf []byte) (*logRecordHeader, int64) {
	return nil, 0
}

func getLogRecordCRC(lr *LogRecord, header []byte) uint32 {
	return 0
}

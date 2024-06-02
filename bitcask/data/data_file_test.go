package data

import (
	"github.com/qixia1998/GoKVStroge/bitcask/fio"
	"os"
	"reflect"
	"testing"
)

func TestDataFile_Close(t *testing.T) {
	type fields struct {
		FileId    uint32
		WriteOff  int64
		IoManager fio.IOManager
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &DataFile{
				FileId:    tt.fields.FileId,
				WriteOff:  tt.fields.WriteOff,
				IoManager: tt.fields.IoManager,
			}
			if err := df.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataFile_ReadLogRecord(t *testing.T) {
	type fields struct {
		FileId    uint32
		WriteOff  int64
		IoManager fio.IOManager
	}
	type args struct {
		offset int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LogRecord
		want1   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &DataFile{
				FileId:    tt.fields.FileId,
				WriteOff:  tt.fields.WriteOff,
				IoManager: tt.fields.IoManager,
			}
			got, got1, err := df.ReadLogRecord(tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLogRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLogRecord() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReadLogRecord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDataFile_Sync(t *testing.T) {
	type fields struct {
		FileId    uint32
		WriteOff  int64
		IoManager fio.IOManager
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &DataFile{
				FileId:    tt.fields.FileId,
				WriteOff:  tt.fields.WriteOff,
				IoManager: tt.fields.IoManager,
			}
			if err := df.Sync(); (err != nil) != tt.wantErr {
				t.Errorf("Sync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataFile_Write(t *testing.T) {
	type fields struct {
		FileId    uint32
		WriteOff  int64
		IoManager fio.IOManager
	}
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &DataFile{
				FileId:    tt.fields.FileId,
				WriteOff:  tt.fields.WriteOff,
				IoManager: tt.fields.IoManager,
			}
			if err := df.Write(tt.args.buf); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataFile_readNBytes(t *testing.T) {
	type fields struct {
		FileId    uint32
		WriteOff  int64
		IoManager fio.IOManager
	}
	type args struct {
		n      int64
		offset int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantB   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &DataFile{
				FileId:    tt.fields.FileId,
				WriteOff:  tt.fields.WriteOff,
				IoManager: tt.fields.IoManager,
			}
			gotB, err := df.readNBytes(tt.args.n, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("readNBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("readNBytes() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestOpenDataFile(t *testing.T) {
	dataFile1, err := OpenDataFile(os.TempDir(), 0)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile1)

}

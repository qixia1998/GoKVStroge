package fio

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestFileIO_Close(t *testing.T) {

}

func TestFileIO_Read(t *testing.T) {

}

func TestFileIO_Sync(t *testing.T) {

}

func TestFileIO_Write(t *testing.T) {

}

func TestNewFileIOManager(t *testing.T) {
	fio, err := NewFileIOManager(filepath.Join("/tmp", "1.data"))
	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

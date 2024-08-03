package classpath

import (
	"os"
	"path/filepath"
)

// DirEntry
// 目录形式的类路径
// /**
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	// 将参数转化成绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (e *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// TODO 这里为什么要把Entry本身也返回回来?
	fileName := filepath.Join(e.absDir, className)
	data, err := os.ReadFile(fileName)
	return data, e, err
}

func (e *DirEntry) String() string {
	return e.absDir
}

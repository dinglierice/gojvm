package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
}

func newEntryZip(path string) *ZipEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{abs}
}

func (e *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(className)
	if err != nil {
		return nil, nil, err
	}

	defer reader.Close()
	for _, f := range reader.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			// TODO defer放在for循环内可能会产生潜在的问题
			defer rc.Close()
			readAll, err := io.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return readAll, e, nil
		}
	}

	return nil, nil, errors.New("class not found : " + className)
}

func (e *ZipEntry) String() string {
	return e.absDir
}

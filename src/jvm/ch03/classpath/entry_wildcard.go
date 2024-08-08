package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// 创建一个通配符入口
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	var compositeEntry []Entry

	waldFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	err := filepath.Walk(baseDir, waldFn)
	if err != nil {
		return nil
	}
	return compositeEntry
}

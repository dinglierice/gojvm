package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(classname string) ([]byte, Entry, error)
	String() string
}

/*
*根据参数创建不同类型的Entry
 */
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") ||
		strings.HasSuffix(path, "zip") || strings.HasSuffix(path, "ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

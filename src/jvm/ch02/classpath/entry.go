package classpath

import (
	"os"
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
	return nil
}

package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (e CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range e {
		data, et, err := entry.readClass(className)
		if err == nil {
			return data, et, err
		}
	}
	return nil, nil, errors.New("class not found : " + className)
}

func (e CompositeEntry) String() string {
	strs := make([]string, len(e))
	for i, entry := range e {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}

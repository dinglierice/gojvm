package classfile

import "encoding/binary"

// ClassReader /*
type ClassReader struct {
	data []byte
}

func (e *ClassReader) readUint8() uint8 {
	val := e.data[0]
	e.data = e.data[1:]
	return val
}

func (e *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(e.data)
	e.data = e.data[2:]
	return val
}

func (e *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(e.data)
	e.data = e.data[4:]
	return val
}

func (e *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(e.data)
	e.data = e.data[8:]
	return val
}

func (e *ClassReader) readUint16s() []uint16 {
	n := e.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = e.readUint16()
	}
	return s
}

func (e *ClassReader) readBytes(length uint32) []byte {
	bytes := e.data[:length]
	e.data = e.data[length:]
	return bytes
}

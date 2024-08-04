package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (e *ConstantStringInfo) readInfo(reader *ClassReader) {
	e.stringIndex = reader.readUint16()
}

func (e *ConstantStringInfo) String() string {
	return e.cp.getUtf8(e.stringIndex)
}

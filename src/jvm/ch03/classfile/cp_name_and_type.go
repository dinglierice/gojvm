package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex uint16
	typeIndex uint16
}

func (e *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	e.nameIndex = reader.readUint16()
	e.typeIndex = reader.readUint16()
}

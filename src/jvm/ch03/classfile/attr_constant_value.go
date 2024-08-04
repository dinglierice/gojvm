package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (e *ConstantValueAttribute) readInfo(reader *ClassReader) {
	e.constantValueIndex = reader.readUint16()
}

func (e *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return e.constantValueIndex
}

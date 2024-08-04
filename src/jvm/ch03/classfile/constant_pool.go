package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (e ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := e[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid ConstantPool Index")
}

func (e ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := e.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := e.getUtf8(ntInfo.nameIndex)
	_type := e.getUtf8(ntInfo.typeIndex)
	return name, _type
}

func (e ConstantPool) getClassName(index uint16) string {
	classInfo := e.getConstantInfo(index).(*ConstantClassInfo)
	return e.getUtf8(classInfo.nameIndex)
}

func (e ConstantPool) getUtf8(index uint16) string {
	utf8Info := e.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

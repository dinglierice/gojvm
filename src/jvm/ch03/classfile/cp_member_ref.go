package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (e *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	e.classIndex = reader.readUint16()
	e.nameAndTypeIndex = reader.readUint16()
}

func (e *ConstantMemberrefInfo) ClassName() string {
	return e.cp.getClassName(e.classIndex)
}

func (e *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return e.cp.getNameAndType(e.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

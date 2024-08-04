package classfile

const (
	Constant_Class              = 7
	Constant_Fieldref           = 9
	Constant_Methodref          = 10
	Constant_InterfaceMethodref = 11
	Constant_String             = 8
	Constant_Integer            = 3
	Constant_Float              = 4
	Constant_Long               = 5
	Constant_Double             = 6
	Constant_NameAndType        = 12
	Constant_Utf8               = 1
	Constant_MethodHandle       = 15
	Constant_MethodType         = 16
	Constant_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case Constant_Integer:
		return &ConstantIntegerInfo{}
	case Constant_Float:
		return &ConstantFloatInfo{}
	case Constant_Long:
		return &ConstantLongInfo{}
	case Constant_Double:
		return &ConstantDoubleInfo{}
	case Constant_Utf8:
		return &ConstantUtf8Info{}
	case Constant_String:
		return &ConstantStringInfo{cp: cp}
	case Constant_Class:
		return &ConstantClassInfo{cp: cp}
	case Constant_Fieldref:
		return &ConstantFieldrefInfo{
			ConstantMemberrefInfo{
				cp: cp,
			},
		}
	case Constant_Methodref:
		return &ConstantMethodrefInfo{
			ConstantMemberrefInfo{
				cp: cp,
			},
		}
	case Constant_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{
			ConstantMemberrefInfo{
				cp: cp,
			},
		}
	case Constant_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case Constant_MethodType:
		return &ConstantMethodTypeInfo{}
	case Constant_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case Constant_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag")
	}
}

package classfile

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (e *MarkerAttribute) readInfo(reader *ClassReader) {
	// 不读取任何信息
}

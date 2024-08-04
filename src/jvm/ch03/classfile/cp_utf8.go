package classfile

type ConstantUtf8Info struct {
	str string
}

// TODO 这里存在潜在的问题
// JAVA 的UTF8采用的编码解码方式非标准UTF-8
func (e *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	e.str = string(bytes)
}

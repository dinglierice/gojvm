package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (e *SourceFileAttribute) readInfo(reader *ClassReader) {
	e.sourceFileIndex = reader.readUint16()
}

func (e *SourceFileAttribute) FileName() string {
	return e.cp.getUtf8(e.sourceFileIndex)
}

package classfile

import (
	"errors"
	"fmt"
)

func (p ConstantPool) getClassName(class uint16) string {
	return ""
}

type AttributeInfo struct {
}

type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	// 类访问标志
	accessFlags uint16

	// 类、超类以及接口的地址索引
	thisClass  uint16
	superClass uint16
	interfaces []uint16
	fields     []*MemberInfo
	methods    []*MemberInfo
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = errors.New(fmt.Sprint("panic: %v", r))
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/*
*
检测文件类型
*/
func (e *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
*
检测文件版本，默认向前兼容，这里只检测java8的标识
[注意这里的readUnit16方法，实际每次调用都会改变包装类持有的切片对象[bytes]]
*/
func (e *ClassFile) readAndCheckVersion(reader *ClassReader) {
	e.minorVersion = reader.readUint16()
	e.majorVersion = reader.readUint16()
	switch e.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if e.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError")
}

func (e *ClassFile) MinorVersion() uint16 {
	return e.minorVersion
}

func (e *ClassFile) MajorVersion() uint16 {
	return e.majorVersion
}

func (e *ClassFile) ConstantPool() ConstantPool {
	return e.constantPool
}

func (e *ClassFile) AccessFlags() uint16 {
	return e.accessFlags
}

func (e *ClassFile) Fields() []*MemberInfo {
	return e.fields // 假设self.fields存储了实际的值
}

func (e *ClassFile) Methods() []*MemberInfo {
	return e.methods // 假设self.methods存储了实际的值
}

func (e *ClassFile) ClassName() string {
	return e.constantPool.getClassName(e.thisClass)
}

func (e *ClassFile) SuperClassName() string {
	if e.superClass > 0 {
		return e.constantPool.getClassName(e.superClass)
	}
	return ""
}

func (e *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(e.interfaces))
	for i, cpIndex := range e.interfaces {
		interfaceNames[i] = e.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (e *ClassFile) read(reader *ClassReader) {
	e.readAndCheckMagic(reader)
	e.readAndCheckVersion(reader)
	e.constantPool = readConstantPool()
	e.accessFlags = reader.readUint16()
	e.thisClass = reader.readUint16()
	e.superClass = reader.readUint16()
	e.interfaces = reader.readUint16s()
	e.fields = readMembers(reader, e.constantPool)
	e.methods = readMembers(reader, e.constantPool)
	e.attributes = readAttributes(reader, e.constantPool)
}

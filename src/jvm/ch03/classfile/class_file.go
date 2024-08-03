package classfile

import (
	"errors"
	"fmt"
)

type ConstantPool struct {
}

type MemberInfo struct {
}

type AttributeInfo struct {
}

type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
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

	cl := &ClassReader{classData}
	cf := &ClassFile{}
	cf.read(cr)
	return
}

func (e *ClassFile) readAndCheckMagic(reader *ClassReader) {
	// 默认实现，假设不会执行任何操作
}

func (e *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 默认实现，假设不会执行任何操作
}

func (e *ClassFile) MinorVersion() uint16 {
	return e.minorVersion // 假设self.minorVersion存储了实际的值
}

func (e *ClassFile) MajorVersion() uint16 {
	return e.majorVersion // 假设self.majorVersion存储了实际的值
}

func (e *ClassFile) ConstantPool() ConstantPool {
	return e.constantPool // 假设self.constantPool存储了实际的值
}

func (e *ClassFile) AccessFlags() uint16 {
	return e.accessFlags // 假设self.accessFlags存储了实际的值
}

func (e *ClassFile) Fields() []*MemberInfo {
	return e.fields // 假设self.fields存储了实际的值
}

func (e *ClassFile) Methods() []*MemberInfo {
	return e.methods // 假设self.methods存储了实际的值
}

func (e *ClassFile) ClassName() string {
	return e.ClassName() // TODO 待实现
}

func (e *ClassFile) SuperClassName() string {
	return e.SuperClassName() // TODO 待实现
}

func (e *ClassFile) InterfaceNames() []string {
	return e.InterfaceNames() // TODO 待实现
}

func (e *ClassFile) read(reader *ClassReader) {
	e.readAndCheckMagic(reader)
	e.readAndCheckVersion(reader)
	e.constantPool = readConstantPool()
}

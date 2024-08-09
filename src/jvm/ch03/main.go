package main

import (
	"fmt"
	"strings"
	"zydl.edu/start/ch03/classfile"
	"zydl.edu/start/ch03/classpath"
)

// JVM启动流程：
// 解析命令行参数
// 处理版本查询和帮助请求
// 设置类路径
// 加载指定的主类
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	// 类路径是 JVM 或 Java 编译器用来查找类（class）和资源文件的路径。它告诉 Java 运行时环境去哪里查找类和包。
	// jre类路径: runtime + 标准库。jre 通常不需要显式指定
	// cpOption: 用户自己的类和第三方库的路径

	// 返回ClassPath类型的引用
	// ClassPath持有Entry
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Printf("Load class:%s successfully\n", cmd.class)
	printClassInfo(cf)

	class, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Read class %s error:%v", className, err)
		return
	}

	// 正确加载了启动类
	fmt.Printf("class data : %v\n", class)
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Println("Class Info:")
	fmt.Printf("  Version: %d.%d\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("  constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("  Access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("  This class: %s\n", cf.ClassName())
	fmt.Printf("  Super class: %s\n", cf.SuperClassName())
	fmt.Println("  Interfaces:")
	fmt.Println("  Fields:")
	for _, f := range cf.Fields() {
		fmt.Printf("    %s\n", f.Name())
	}
	fmt.Println("  Methods:")
	for _, m := range cf.Methods() {
		fmt.Printf("    %s\n", m.Name())
	}
}

func loadClass(name string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(name)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

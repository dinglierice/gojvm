package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (e *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	// 每一种类路径entry对应着不同的接口实现，所以readClass能够找到对应的实现
	// 变量初始化必须是接口的具体实现
	// 接口的零值是 nil，直接调用 nil 接口的方法会导致运行时 panic。
	if data, entry, err := e.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := e.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return e.userClasspath.readClass(className)
}

func (e *Classpath) String() string {
	return e.userClasspath.String()
}

// 搜索
// 仅仅是为了检验文件是否存在，没输入或者不存在则输入系统环境变量
// bootClasspath Entry 引导类路径：核心库和基础库
// extClasspath  Entry 扩展类路径：基础库扩展
func (e *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	e.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	e.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("can not find a jre folder")
}

func exists(option string) bool {
	if _, err := os.Stat(option); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (e *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	e.userClasspath = newEntry(cpOption)
}

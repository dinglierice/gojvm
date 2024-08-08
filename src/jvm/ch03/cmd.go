package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	XjreOption  string
	cpOption    string
	class       string
	args        []string
}

// & 操作符（取地址）：
// & 用于获取一个变量的内存地址。
// 它可以用于任何变量前面，返回该变量的指针。
// 例如：ptr := &x 意味着 ptr 现在存储了变量 x 的内存地址。
// * 操作符（解引用）：
// * 有两个用途：
// a. 当用在类型前面时，它用于声明指针类型。
// b. 当用在指针变量前面时，它用于解引用（获取指针指向的值）。
// 例如：
// var ptr *int 声明 ptr 为整数指针类型。
// value := *ptr 获取 ptr 指向的整数值。
func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version information and exit")
	flag.StringVar(&cmd.cpOption, "classPath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: ·%s [-options] class [args...]\n", os.Args[0])
}

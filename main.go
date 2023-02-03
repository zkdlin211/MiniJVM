package main

import (
	"MiniJVM/src/classfile"
	"MiniJVM/src/classpath"
	"fmt"
	"strings"
)

var CMD *Cmd

func main() {
	CMD = parseCmd()
	if CMD.versionFlag {
		fmt.Println("version 0.0.1")
	} else if CMD.helpFlag || CMD.class == "" {
		printUsage()
	} else {
		startJVM(CMD)
	}
}

func startJVM(cmd *Cmd) {
	classPath := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(cmd, classPath, className)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func loadClass(cmd *Cmd, classPath *classpath.Classpath, className string) *classfile.ClassFile {
	classData, _, err := classPath.ReadClass(className)
	if err != nil {
		fmt.Printf("Error finding or loading main class %s\n", cmd.class)
		fmt.Printf(err.Error())
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}

//func startJVM(CMD *Cmd) {
//	classPath := classpath.Parse(CMD.XjreOption, CMD.cpOption)
//	fmt.Printf("classpath:%s class:%s args:%v\n",
//		classPath, CMD.class, CMD.args)
//	className := strings.Replace(CMD.class, ".", "/", -1)
//	classData, _, err := classPath.ReadClass(className)
//	if err != nil {
//		fmt.Printf("Error finding or loading main class %s\n", CMD.class)
//		fmt.Printf(err.Error())
//		return
//	}
//	fmt.Printf("class data:%v\n", classData)
//}

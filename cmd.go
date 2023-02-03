package main

import (
	"MiniJVM/src/runtime"
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string //classpath option
	XjreOption  string // -Xjre,specify the location of the jre directory
	XmsOption   string // -Xms: Sets the initial size of the heap.
	XmxOption   string // -Xmx: Sets the maximum size of the heap.
	XssOption   string // -Xss: Sets the maximum size of the stack.
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classPath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.StringVar(&cmd.XmsOption, "Xms", runtime.XmsOption, "initial size of the heap")
	flag.StringVar(&cmd.XmxOption, "Xmx", runtime.Xmx, "maximum size of the heap")
	flag.StringVar(&cmd.XssOption, "Xss", runtime.XssOption, "maximum size of the stack")
	flag.StringVar(&cmd.cpOption, "cp", "", "classPath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: &s [-options] class [args...]\n", os.Args)
}

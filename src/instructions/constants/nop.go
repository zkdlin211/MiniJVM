package constants

import (
	"MiniJVM/src/instructions/base"
	"MiniJVM/src/runtime"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *runtime.Frame) {
	// do nothing
}

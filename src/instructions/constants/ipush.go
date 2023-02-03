package constants

import (
	"MiniJVM/src/instructions/base"
	"MiniJVM/src/runtime"
)

// BIPUSH : The bipush instruction takes a byte integer from the operand,
// expands it to an int, and then pushes it to the top of the stack
type BIPUSH struct {
	// push byte
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *runtime.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// SIPUSH : The sipush instruction takes a short integer from the operand,
// expands it to an int, and pushes it to the top of the stack.
type SIPUSH struct {
	// push short
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *runtime.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

package constants

import (
	"MiniJVM/src/instructions/base"
	"MiniJVM/src/runtime"
)

// ACONST_NULL instruction pushes the `null` reference to the top of the operand stack
type ACONST_NULL struct{ base.NoOperandsInstruction }

func (self *ACONST_NULL) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushRef(nil)
}

// DCONST_0 instruction pushes double type0 to the top of the operand stack
type DCONST_0 struct{ base.NoOperandsInstruction }

func (self *DCONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOperandsInstruction }

func (self *DCONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

type FCONST_0 struct{ base.NoOperandsInstruction }

func (self *FCONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOperandsInstruction }

func (self *FCONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct{ base.NoOperandsInstruction }

func (self *FCONST_2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// ICONST_M1 instruction pushes int type -1 to the top of the operand stack
type ICONST_M1 struct{ base.NoOperandsInstruction }

func (self *ICONST_M1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct{ base.NoOperandsInstruction }

func (self *ICONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct{ base.NoOperandsInstruction }

func (self *ICONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct{ base.NoOperandsInstruction }

func (self *ICONST_2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct{ base.NoOperandsInstruction }

func (self *ICONST_3) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }

func (self *ICONST_4) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct{ base.NoOperandsInstruction }

func (self *ICONST_5) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(5)
}

type LCONST_0 struct{ base.NoOperandsInstruction }

func (self *LCONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }

func (self *LCONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(1)
}

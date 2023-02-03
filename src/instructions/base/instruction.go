package base

import "MiniJVM/src/runtime"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *runtime.Frame)
}

// NoOperandsInstruction represents an instruction that has no operands, so no fields are defined.
type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// BranchInstruction represents the jump instruction, and the Offset field stores the jump offset.
type BranchInstruction struct {
	Offset int
}

// FetchOperands method reads an uint16 integer from the bytecode,
// which uint16 converted to int and uint16 assigned to the Offset field.
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// Index8Instruction is an abstraction of load and store type instructions,
// index field indicates the index of the local variable table
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// Index16Instruction is an abstraction of some instructions that require access to the run-time constant pool,
// whose index is given by a 2-byte parameter
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}

package runtime

// Frame is a data structure that represents a method call, including its local variables, operand stack,
// and other information. When a method is called, a new frame is pushed onto the stack, and when a method returns,
// its frame is popped from the stack.
type Frame struct {
	lower *Frame
	// The local table size and operand stack depth required to execute the method are pre-calculated by the compiler
	// and stored in the Code attribute of the class file method_info structure
	localVars   LocalVars
	operandSack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:   newLocalVars(maxLocals),
		operandSack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandSack
}

package runtime

type Thread struct {
	pc int

	// Stack Pointer of the Java virtual machine stack
	stack *Stack
}

func newThread() *Thread {

	return &Thread{

		// todo
		stack: newStack(1024),
	}
}

func (self *Thread) getPC() int {
	return self.pc
}

func (self *Thread) setPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.peek()
}

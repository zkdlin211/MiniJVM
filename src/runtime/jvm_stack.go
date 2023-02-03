package runtime

type Stack struct {
	// the capacity of the stack (how many frames can be held at most)
	maxSize uint
	// current size of the stack
	size uint
	// peek frame pointer of the stack
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// pushing a new frame onto the stack
// throw an error when exceeding max stack size
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

// popping the top frame from the stack
func (self *Stack) pop() (frame *Frame) {
	if self.isEmpty() {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	// todo: gc
	top.lower = nil
	self.size--
	return top
}

// accessing the top frame on the stack
func (self *Stack) peek() (frame *Frame) {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}

func (self *Stack) isEmpty() bool {
	return self.size == 0
}

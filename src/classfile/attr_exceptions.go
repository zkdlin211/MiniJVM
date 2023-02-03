package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16Table()
}

func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}

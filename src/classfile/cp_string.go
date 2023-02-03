package classfile

// ConstantStringInfo represents the java.lang.String literal
type ConstantStringInfo struct {
	cp ConstantPool

	//stringIndex is used to read the reference from the constant pool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// String method loops up UTF-8 string from constant pool by given stringIndex
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}

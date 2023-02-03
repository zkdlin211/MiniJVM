package classfile

// ConstantValueAttribute represents a fixed-length property that occurs only in the field_info structure
// and is used to represent the value of a constant expression
// (see Section 15.28 of the Java Language specification for details).
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}

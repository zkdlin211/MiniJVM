package classfile

// ConstantNameAndTypeInfo gives the name and descriptor of a field or method.
// ConstantClassInfo and ConstantNameAndTypeInfo together can uniquely identify a field or method.
// The field or method name is given by name_index, and the descriptor of the field or method is given by
// descriptor_index. name_index and descriptor_index are both constant pool references to the ConstantUtf8Info constant.
// The field and method names are the names of the fields or methods that appear in the code
// (or are generated by the compiler).
// See the Java Virtual Machine specification for definitions.
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}

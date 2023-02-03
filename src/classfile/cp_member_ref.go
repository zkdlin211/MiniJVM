package classfile

// ConstantMemberrefInfo is a general implementation representing
// ConstantFieldrefInfo, ConstantMethodrefInfo, and ConstantInterfaceMethodrefInfo.
// ConstantFieldrefInfo represents field symbolic references,
// ConstantMethodrefInfo represents normal (non-interface) method symbolic references, and
// ConstantInterfaceMethodrefInfo represents interface method symbolic references.
// class_index and name_and_type_index are both constant pool indexes that point to
// ConstantClassInfo and ConstantNameAndTypeInfo constants, respectively.
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

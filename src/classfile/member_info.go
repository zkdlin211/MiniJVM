package classfile

// MemberInfo structure is used by the Java Virtual Machine (JVM)
// to load the class file into memory and to execute the methods and
// access the fields defined in the class.
// The JVM uses the information in the memberInfo structures to create
// the appropriate objects in memory and to manage access to the fields and methods of the class.
// A Java class file contains several memberInfo structures,
// each of which represents a single field or method in the class.
// The MemberInfo structure includes information such as the name and type of the member,
// access flags, and attributes such as code for methods and constant values for fields.
type MemberInfo struct {
	constantPool    ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, constantPool ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, constantPool)
	}
	return members
}

func readMember(reader *ClassReader, constantPool ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    constantPool,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, constantPool),
	}
}

// AccessFlags is a Getter of accessFlags in MemberInfo
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

// Name gets field name or method name from constant pool by nameIndex
func (self *MemberInfo) Name() string {
	return self.constantPool.getUtf8(self.nameIndex)
}

// Descriptor gets field or method descriptor from constant pool by descriptorIndex
func (self *MemberInfo) Descriptor() string {
	return self.constantPool.getUtf8(self.descriptorIndex)
}

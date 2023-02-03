package classfile

import "fmt"

// The ClassFile structure exactly mirrors the class file schema
// defined by the Java Virtual Machine specification.
// Due to the nature of the Go language, this ClassFile structure has been modified from the original specification,
// noting that constant_pool_count, interfaces_count, fields_count, methods_count have been omitted.
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	cp           ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse methods parses []byte into a ClassFile entity
func Parse(classData []byte) (classFile *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	classReader := &ClassReader{classData}
	classFile = &ClassFile{}
	classFile.read(classReader)
	return
}

// parse data byte[] in reader to classFile
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.cp = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16Table()
	self.fields = readMembers(reader, self.cp)
	self.methods = readMembers(reader, self.cp)
	self.attributes = readAttributes(reader, self.cp)
}

// readAndCheckMagic checks if this class file starts with 0xCAFEBABE
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	self.magic = reader.readUint32()
	if self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: Incompatible magic value in class file!")
	}
}

// Refer to Java 8, class files with version 45.0 to 52.0 are supported
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// MajorVersion is a Getter of majorVersion in ClassFile
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

// MinorVersion is a Getter of minorVersion in ClassFile
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

// ConstantPool is a Getter of cp in ClassFile
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.cp
}

// AccessFlags is a Getter of accessFlags in ClassFile
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

// Fields is a Getter of fields in ClassFile
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

// Methods is a Getter of methods in ClassFile
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

// ClassName gets this className from constant pool
func (self *ClassFile) ClassName() string {
	return self.cp.getClassName(self.thisClass)
}

// superClassName gets super className from constant pool
func (self *ClassFile) superClassName() string {
	if self.superClass > 0 {
		return self.cp.getClassName(self.superClass)
	}
	// only java.lang.Object does not have a super class
	return ""
}

// InterfaceNames gets all interface name(s) from constant pool
func (self *ClassFile) InterfaceNames() []string {
	interfacesNames := make([]string, len(self.interfaces))
	for i, constantPoolIndex := range self.interfaces {
		interfacesNames[i] = self.cp.getClassName(constantPoolIndex)
	}
	return interfacesNames
}

// SuperClassName gets super class name of current class
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.cp.getClassName(self.superClass)
	}
	return ""
}

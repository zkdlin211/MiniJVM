package classfile

// ConstantPool is a table of constants used in a Java class file. It is used to store constant values
// such as string literals, method names, field names, and other constant values used in the class.
// The constant pool is stored in the class file as an array of constant_info structures.
// The JVM uses the constant pool to resolve references to constant values, such as string literals, at runtime.
// The size of the constant pool given in the header is one larger than it really is.
// valid index range in ConstantPool is 1 to n-1, index 0 is invalid and not used for pointing any constant.
// CONSTANT_Long_info and CONSTANT_Double_info occupy two spaces. If these two constants are present
// in the constant pool, the actual number of constants is less than n-1.
// Also, some numbers from 1 to n-1 will become invalid indexes.
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	constantPools := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		constantPools[i] = readConstantInfo(reader, constantPools)
		switch constantPools[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return constantPools
}

// getConstantInfo method gets a specific ConstantInfo by index
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// getNameAndType method looks up the name and descriptor of a field or method from constant pool
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	nameTypeInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(nameTypeInfo.nameIndex)
	_type := self.getUtf8(nameTypeInfo.descriptorIndex)
	return name, _type
}

// getClassName method loops up class name from constant pool
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

// getUtf8 method loops up UTF-8 string from constant pool
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

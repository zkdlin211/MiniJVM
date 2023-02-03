package classfile

/*
	SourceFile_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 sourcefile_index;
	}

SourceFile is a long selectable attribute that appears only in the ClassFile structure
and is used to indicate the source file name.
*/
type SourceFileAttribute struct {
	cp ConstantPool
	// sourceFileIndex represents constant pool index, and points to the ConstantUtf8Info constant
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}
func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}

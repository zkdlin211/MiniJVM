package classfile

/*
	Deprecated_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
	}

	Synthetic_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
	}

Deprecated and Synthetic attributes serve only as a marker and do not contain any data.
*/
type MarkerAttribute struct {
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

package classfile

import "encoding/binary"

// ClassReader is a helper class that provides util methods for manipulating bytes directly
type ClassReader struct {
	data []byte
}

// readUint8 reads the u1 type data
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// readUint16 reads the u2 type data
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// readUint16Table reads the table with uint16 length,
// the size of the table is indicated by the data(uint16) at the beginning
func (self *ClassReader) readUint16Table() []uint16 {
	length := self.readUint16()
	table := make([]uint16, length)
	for i := range table {
		table[i] = self.readUint16()
	}
	return table
}

// readUint32 reads the u4 type data
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// readBytes: reads the specified number of bytes
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}

// readUint64 reads uint64 type data
// notice: u8 is not defined in the JVM specification
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

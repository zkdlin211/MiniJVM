package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

/*
A classpath in the form of a ZIP or JAR file
*/
type ZipEntry struct {
	absPath       string
	zipReadCloser *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath, nil}
}

// extract class files from zip file
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if self.zipReadCloser == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}
	classFile := self.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}
	data, err := readClass(classFile)
	return data, self, err
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	//read class data
	data, err := io.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (self *ZipEntry) openJar() error {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return err
	}
	self.zipReadCloser = r
	return nil
}

func (self *ZipEntry) String() string {
	return self.absPath
}

func (self *ZipEntry) findClass(className string) *zip.File {
	for _, f := range self.zipReadCloser.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

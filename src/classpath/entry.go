package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if isJar(path) ||
		strings.HasSuffix(path, ",zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

func isJar(path string) bool {
	return strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR")
}

package classpath

import (
	"io/fs"
	"path/filepath"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove last "*"
	compositeEnry := []Entry{}
	err := filepath.Walk(baseDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if isJar(path) {
			compositeEnry = append(compositeEnry, newZipEntry(path))
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return compositeEnry
}

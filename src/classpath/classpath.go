package classpath

import (
	"os"
	"path/filepath"
)

//searches for classes based on class paths. In the order of search, they can be divided into
//
//bootstrap classpath (bootstrap classpath)
//Extension classpath (extension classpath)
//User class path (user classpath)
//
//The bootstrap classpath is by default in the jre\lib directory,
//where the java standard libraries (mostly in rt.jar) are located.
//
//The extension class path corresponds to the jre\lib\ext directory,
//and classes that use the java extension mechanism are located in this path.
//
//Our own implemented classes, as well as third-party libraries,
//are located in the user class path. The default value for the user class path is
//the current directory, which is "."

type Classpath struct {
	// Bootstrap classpath
	bootClasspath Entry
	// extension classpat
	extClasspath Entry
	// user classpath
	userClasspath Entry
}

// jreOption:
// usage: "-Xjre" to indicate jre directory.
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// The -Xjre option entered by the user is preferred as the jre directory.
// If this option is not entered, look for the jre directory in the current directory.
// If it is not found, try using the JAVA_HOME environment variable.
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// search current path
	if exists("./jre") {
		return "./jre"
	}
	// JAVA_HOME/jre
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("Cannot find jre folder!")
}

/*
is given path exists
*/
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// If the user does not provide the -classpath/-cp option,
// the current directory is used as the user classpath.
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

// ReadClass searches for class files from the boostrap classpath,
// extension classpath, and user classpath in turn
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

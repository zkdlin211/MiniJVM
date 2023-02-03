package runtime

// Slot is a structure that used to store values in the local variables array of a Java virtual machine (JVM)
// stack frame. The num field is used to store integer values and the ref field is used to store references to objects.
// In the JVM, local variables are used to store intermediate results and method parameters.
// Each method call has its own set of local variables stored in a separate frame on the JVM stack.
type Slot struct {
	num int32
	ref *Object
}

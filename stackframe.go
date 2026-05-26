package errors

import (
	"runtime"
)

// A StackFrame contains all necessary information about to generate a line
// in a callstack.
type StackFrame struct {
	// The path to the file containing this ProgramCounter
	File string
	// The LineNumber in that file
	LineNumber int
	// The Name of the function that contains this ProgramCounter
	Name string
	// The Package that contains this function
	Package string
	// The underlying ProgramCounter
	ProgramCounter uintptr
}

// NewStackFrame popoulates a stack frame object from the program counter.
func NewStackFrame(pc uintptr) (frame StackFrame) {
	_ = "STUB: not implemented"
	return *new(StackFrame)
}

// pc -1 because the program counters we use are usually return addresses,
// and we want to show the line that corresponds to the function call

// Func returns the function that contained this frame.
func (frame *StackFrame) Func() *runtime.Func { _ = "STUB: not implemented"; return nil }

// String returns the stackframe formatted in the same way as go does
// in runtime/debug.Stack()
func (frame *StackFrame) String() string { _ = "STUB: not implemented"; return "" }

// SourceLine gets the line of code (from File and Line) of the original source if possible.
func (frame *StackFrame) SourceLine() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (frame *StackFrame) sourceLine() (string, error) { _ = "STUB: not implemented"; return "", nil }

func packageAndName(fn *runtime.Func) (string, string) {
	_ = "STUB: not implemented"
	return "",

		// The name includes the path name to the package, which is unnecessary
		// since the file name is already included.  Plus, it has center dots.
		// That is, we see
		//
		//	runtime/debug.*T·ptrmethod
		//
		// and want
		//
		//	*T.ptrmethod
		//
		// Since the package path might contains dots (e.g. code.google.com/...),
		// we first remove the path prefix if there is one.
		""
}

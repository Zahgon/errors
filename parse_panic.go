package errors

type uncaughtPanic struct{ message string }

func (p uncaughtPanic) Error() string {
	_ = "STUB: not implemented"

	// ParsePanic allows you to get an error object from the output of a go program
	// that panicked. This is particularly useful with https://github.com/mitchellh/panicwrap.
	return ""
}

func ParsePanic(text string) (*Error, error) { _ = "STUB: not implemented"; return nil, nil }

// The lines we're passing look like this:
//
//	main.(*foo).destruct(0xc208067e98)
//	        /0/go/src/github.com/bugsnag/bugsnag-go/pan/main.go:22 +0x151
func parsePanicFrame(name string, line string, createdBy bool) (*StackFrame, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

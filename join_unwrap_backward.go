//go:build !go1.20
// +build !go1.20

package errors

// Disclaimer: functions Join and Unwrap are copied from the stdlib errors
// package v1.21.0.

// Join returns an error that wraps the given errors.
// Any nil error values are discarded.
// Join returns nil if every value in errs is nil.
// The error formats as the concatenation of the strings obtained
// by calling the Error method of each element of errs, with a newline
// between each string.
//
// A non-nil error returned by Join implements the Unwrap() []error method.
func Join(errs ...error) error { _ = "STUB: not implemented"; return nil }

type joinError struct {
	errs []error
}

func (e *joinError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *joinError) Unwrap() []error {
	_ = "STUB: not implemented"

	// Unwrap returns the result of calling the Unwrap method on err, if err's
	// type contains an Unwrap method returning error.
	// Otherwise, Unwrap returns nil.
	//
	// Unwrap only calls a method of the form "Unwrap() error".
	// In particular Unwrap does not unwrap errors returned by [Join].
	return nil
}

func Unwrap(err error) error { _ = "STUB: not implemented"; return nil }

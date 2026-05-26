//go:build go1.13
// +build go1.13

package errors

// As finds the first error in err's tree that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
//
// For more information see stdlib errors.As.
func As(err error, target interface{}) bool { _ = "STUB: not implemented"; return false }

// Is detects whether the error is equal to a given error. Errors
// are considered equal by this function if they are matched by errors.Is
// or if their contained errors are matched through errors.Is.
func Is(e error, original error) bool { _ = "STUB: not implemented"; return false }

package cef

import (
	"fmt"
)

// CodeError wraps a error code (int) into a Go error
type CodeError struct {
	Code        int
	ErrorFormat string
}

// Error returns the error string for this error
func (ce *CodeError) Error() string {
	return fmt.Errorf(ce.ErrorFormat, ce.Code)
}

package checker

import "fmt"

type UnreachableError struct {
	Path string
	Err  error
}

func (e *UnreachableError) Error() string {
	return fmt.Sprintf("File not found: %s (%v)", e.Path, e.Err)
}

func (e *UnreachableError) Unwrap() error {
	return e.Err
}

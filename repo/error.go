package repo

import "errors"

var (
	ErrNotFound       = errors.New("resource not found")
	ErrNotImplemented = errors.New("not implemented")
	ErrOutOfRange     = errors.New("out of range")
)

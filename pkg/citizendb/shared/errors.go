package shared

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrOutOfRange        = errors.New("out of range")
	ErrInvalidDocumentID = errors.New("invalid document id")
)

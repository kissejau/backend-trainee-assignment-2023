package errors

import "fmt"

var (
	ErrReadingBody     = fmt.Errorf("reading Body error")
	ErrInvalidBody     = fmt.Errorf("invalid Body error")
	ErrIncorrectHeader = fmt.Errorf("incorrect or not passed Header error")
	ErrIncorrectParams = fmt.Errorf("incorrect params")
	ErrInvalidParams   = fmt.Errorf("invalid params")
)

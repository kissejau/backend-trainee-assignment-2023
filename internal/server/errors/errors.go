package errors

import "fmt"

var (
	ErrReadingBody     = fmt.Errorf("reading Body error")
	ErrInvalidBody     = fmt.Errorf("invalid Body error")
	ErrIncorrectHeader = fmt.Errorf("incorrect or not passed Header error")
)

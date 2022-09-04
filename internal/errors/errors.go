package errors

import "fmt"

type errorMsg string

var (
	NotSupportRunnerType errorMsg = "not support runner (%s) type"
)

type Error struct {
	Msg string
}

func (e *Error) Error() string {
	return e.Msg
}

func New(msg errorMsg, params ...interface{}) error {
	return &Error{
		Msg: fmt.Sprintf(string(msg), params...),
	}
}

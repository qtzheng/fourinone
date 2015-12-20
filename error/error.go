package error

import (
	"bytes"
)

const (
	System      = 0
	KeyNotFound = 100
	NodeExist   = 101
)

var (
	errors = map[int]string{
		System:      "系统错误",
		KeyNotFound: "无效的KEY",
		NodeExist:   "节点已经存在",
	}
)

type Error struct {
	ErrorCode  int
	Message    string
	Cause      string
	InnerError *Error
}

func (e Error) Error() string {
	buf := new(bytes.Buffer)
	buf.WriteString(e.errorString())
	innerErr := e.InnerError
	for {
		if innerErr == nil {
			break
		}
		buf.WriteString(innerErr.errorString())
		innerErr = innerErr.InnerError
	}
	return buf.String()
}
func (e Error) errorString() string {
	return e.Message + "(" + e.Cause + ")"
}
func NewError(errorCode int, cause string, innerError error) *Error {
	err := &Error{
		ErrorCode: errorCode,
		Cause:     cause,
		Message:   errors[errorCode],
	}
	if innerError != nil {
		if e, ok := innerError.(*Error); ok {
			err.InnerError = e
		} else {
			err.InnerError = &Error{
				ErrorCode: System,
				Message:   errors[System],
			}
		}
	}

	return err
}

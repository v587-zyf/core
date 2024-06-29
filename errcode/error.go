package errcode

import "fmt"

type Error struct {
	Code    int
	Message string
}

func NewError(code int, message ...string) *Error {
	var msg string

	if len(message) > 0 {
		msg = message[0]
	} else {
		if v, ok := defaultErrs[ErrCode(code)]; !ok {
			msg = fmt.Sprintf("未知代号[%d]", code)
		} else {
			msg = v
		}
	}

	err := &Error{
		Code:    code,
		Message: msg,
	}

	return err
}

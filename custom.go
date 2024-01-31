// Package errors
// @file      : custom.go
// @author    : china.gdxs@gmail.com
// @time      : 2023/12/26 10:39
// @Description:
package errors

type CustomErr struct {
	msg  string
	code int
}

func (e CustomErr) Error() string {
	return e.msg
}

func (e CustomErr) Code() int {
	return e.code
}

func NewCustom(msg string) error {
	return &CustomErr{msg: msg}
}

func NewCustomCode(msg string, code int) error {
	return &CustomErr{
		msg:  msg,
		code: code,
	}
}

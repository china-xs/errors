// Package errors
// @file      : custom.go
// @author    : china.gdxs@gmail.com
// @time      : 2023/12/26 10:39
// @Description:
package errors

import "net/http"

type CustomErr struct {
	msg        string
	code       int
	httpStatus int
}

func (e CustomErr) Error() string {
	return e.msg
}

func (e CustomErr) Code() int {
	return e.code
}
func (e CustomErr) HttpStatus() int {
	return e.httpStatus
}

func NewCustom(msg string) error {
	return &CustomErr{msg: msg, httpStatus: http.StatusOK}
}

func NewCustomCode(msg string, code int) error {
	return &CustomErr{
		msg:        msg,
		code:       code,
		httpStatus: http.StatusOK,
	}
}

func Custom(msg string, code int, httpStatus int) error {
	return &CustomErr{
		msg:        msg,
		code:       code,
		httpStatus: httpStatus,
	}
}

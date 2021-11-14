package errcode

import "fmt"

type ErrorCode struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("error: status=%d, code=%s, msg=%s", e.Status, e.Code, e.Msg)
}

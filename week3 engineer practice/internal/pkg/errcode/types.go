package errcode

var (
	ErrUnknown = &ErrorCode{
		Status: 500,
		Code:   "unknown err",
		Msg:    "unknown err",
	}
)

package util

type Error struct {
	HttpCode     int    `json:"-"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func GetNetErrorWithCode(code int, message string, err error) *Error {
	if err != nil {
		println(err.Error())
	}
	return &Error{
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

func GetNetError(message string, err error) *Error {
	if err != nil {
		println(err.Error())
	}
	return &Error{
		ErrorCode:    -1,
		ErrorMessage: message,
	}
}

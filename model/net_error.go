package model

type NetError struct {
	HttpCode     int    `json:"-"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func GetNetErrorWithCode(code int, message string, err error) *NetError {
	if err != nil {
		println(err.Error())
	}
	return &NetError{
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

func GetNetError(message string, err error) *NetError {
	if err != nil {
		println(err.Error())
	}
	return &NetError{
		ErrorCode:    -1,
		ErrorMessage: message,
	}
}

package lib

type FireResponse struct {
	Errno  int
	ErrMsg string
	Data   interface{}
}

func NewSuccessResponse(data interface{}) FireResponse {
	return FireResponse{
		Errno:  200,
		ErrMsg: "success",
		Data:   data,
	}
}

func NewResponse(errno int, errmsg string, data interface{}) FireResponse {
	return FireResponse{
		Errno:  errno,
		ErrMsg: errmsg,
		Data:   data,
	}
}

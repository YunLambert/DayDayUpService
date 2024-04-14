package global

import "net/http"

var codeMapping = map[int]string{
	http.StatusOK:                  "Success",
	http.StatusInternalServerError: "Internal Server Error",
	http.StatusBadRequest:          "Bad Request",
	http.StatusUnauthorized:        "Unauthorized",
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

func SuccessWithData(data any) Response {
	return Response{
		Code:    http.StatusOK,
		Message: codeMapping[http.StatusOK],
		Data:    data,
	}
}

func ErrorWithMsg(code int, data string) Response {
	return Response{
		Code:    code,
		Message: codeMapping[code],
		Data:    data,
	}
}

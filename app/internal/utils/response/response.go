package response

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type EmptyObject struct{}

func Success(status bool, message string) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  nil,
	}
}

func Error(message string, err string) Response {
	spError := strings.Split(err, "\n")
	return Response{
		Status:  false,
		Message: message,
		Errors:  spError,
	}
}

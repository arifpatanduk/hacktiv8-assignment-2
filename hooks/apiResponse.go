package hooks

import (
	"net/http"
)

// struct for standardized API responses
type Response struct {
	Data    interface{} `json:"data"`
	Message any      `json:"message"`
	Code    int         `json:"code"`
}

// standardized error response
func ErrorResponse(message ...any) Response {
	return Response{
		Message: message,
		Code:    http.StatusBadRequest,
		Data:    nil,
	}
}

// standardized success response
func SuccessResponse(data interface{}, message string) Response {
	return Response{
		Message: message,
		Code:    http.StatusOK,
		Data:    data,
	}
}
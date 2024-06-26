package rest_err

import (
	"net/http"
)


type RestErr struct {
	Message string 	`json:"message"`
	Err 	string	`json:"error"`
	Code 	int		`json:"code"`
	Log		[]Log	`json:"log"`
}

type Log struct {
	Field 	string	`json:"field"`
	Message string	`json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, log []Log) *RestErr{
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Log: log,
	}
}

func NewBadRequestValidadeError(message string, log []Log) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusBadRequest,
		Log: log,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "internal server error",
		Code: http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "not found",
		Code: http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "not found",
		Code: http.StatusNotFound,
	}
}
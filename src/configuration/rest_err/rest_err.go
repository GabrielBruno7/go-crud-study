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

func newRestErr(message, err string, code int, log []Log) *RestErr{
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Log: log,
	}
}

func newBadRequestValidadeError(message string, log []Log) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusBadRequest,
		Log: log,
	}
}

func newBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusBadRequest,
	}
}

func newInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "internal server error",
		Code: http.StatusInternalServerError,
	}
}

func newNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "not found",
		Code: http.StatusNotFound,
	}
}

func newForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "not found",
		Code: http.StatusNotFound,
	}
}
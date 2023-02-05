package errors

import "net/http"

func NewResourceNotFoundError(message string) ApplicationError {
	return ApplicationError{
		message:        message,
		code:           "0104",
		httpStatusCode: http.StatusNotFound,
	}
}

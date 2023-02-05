package errors

import "net/http"

func NewBadOperationError(message string) ApplicationError {
	return ApplicationError{
		message:        message,
		code:           "0100",
		httpStatusCode: http.StatusBadRequest,
	}
}

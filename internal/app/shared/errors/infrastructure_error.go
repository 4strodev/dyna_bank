package errors

import "net/http"

func NewInfrastructureError(message string) ApplicationError {
	return ApplicationError{
		message:        message,
		code:           "0200",
		httpStatusCode: http.StatusInternalServerError,
	}
}

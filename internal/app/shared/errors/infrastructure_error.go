package errors

import "net/http"

type InfrastructureError struct {
	message string
	code    string
}

func NewInfrastructureError(message string) InfrastructureError {
	return InfrastructureError{
		message: message,
		code:    "0200",
	}
}

func (self InfrastructureError) Error() string {
	return self.message
}

func (self InfrastructureError) ErrorCode() string {
	return self.code
}

func (selfl InfrastructureError) HTTPStatusCode() int {
	return http.StatusInternalServerError
}

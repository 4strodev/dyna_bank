package errors

import "net/http"

type BadOperationError struct {
	message string
	code    string
}

func NewBadOperationError(message string) BadOperationError {
	return BadOperationError{
		message: message,
		code:    "0100",
	}
}

func (self BadOperationError) Error() string {
	return self.message
}

func (self BadOperationError) ErrorCode() string {
	return self.code
}

func (self BadOperationError) HTTPStatusCode() int {
	return http.StatusBadRequest
}

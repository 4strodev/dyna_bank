package errors

import "net/http"

type ResourceNotFoundError struct {
	message string
	code    string
}

func NewResourceNotFoundError(message string) ResourceNotFoundError {
	return ResourceNotFoundError{
		message: message,
		code:    "0104",
	}
}

func (self ResourceNotFoundError) Error() string {
	return self.message
}

func (self ResourceNotFoundError) ErrorCode() string {
	return self.code
}

func (self ResourceNotFoundError) HTTPStatusCode() int {
	return http.StatusNotFound
}

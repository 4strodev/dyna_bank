package errors

import "fmt"

type ApplicationErrorInterface interface {
	Error() string
	ErrorCode() string
	HTTPStatusCode() int
	Metadata() error
}

type ApplicationError struct {
	message        string
	code           string
	httpStatusCode int
	metadata       error
}

func (self ApplicationError) Error() string {
	message := self.message
	if self.metadata != nil {
		message = fmt.Sprintf("%s: %s", message, self.metadata)
	}
	return message
}

func (self ApplicationError) ErrorCode() string {
	return self.code
}

func (self ApplicationError) HTTPStatusCode() int {
	return self.httpStatusCode
}

func (self ApplicationError) Metadata() error {
	return self.metadata
}

func (self ApplicationError) WithMetadata(metadata error) ApplicationError {
	self.metadata = metadata
	return self
}

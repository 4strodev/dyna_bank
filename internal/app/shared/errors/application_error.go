package errors

type ApplicationError interface {
	Error() string
	ErrorCode() string
	HTTPStatusCode() int
}

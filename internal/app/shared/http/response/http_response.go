package response

import "github.com/4strodev/dyna-bank/internal/app/shared/http/status"

type HttpResponse struct {
	Status    status.HTTPStatus `json:"status"`
	Message   string            `json:"message,omitempty"`
	ErrorCode string            `json:"error_code,omitempty"`
	Data      any               `json:"data,omitempty"`
}

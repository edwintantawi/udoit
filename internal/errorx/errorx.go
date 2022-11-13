package errorx

import (
	"net/http"
)

type E struct {
	Code    int
	Message string
}

func (e *E) Error() string {
	return e.Message
}

func NewInvariant(m string) *E {
	return &E{Code: http.StatusBadRequest, Message: m}
}

func NewNotFound(m string) *E {
	return &E{Code: http.StatusNotFound, Message: m}
}

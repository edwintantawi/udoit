package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edwintantawi/udoit/internal/errorx"
	"github.com/go-playground/validator/v10"
)

type success struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type fail struct {
	StatusCode int      `json:"code"`
	Message    string   `json:"message"`
	Errors     []string `json:"errors"`
}

type Headers map[string]string

type response struct {
	w http.ResponseWriter
}

func New(w http.ResponseWriter) *response {
	return &response{w: w}
}

func (r *response) encode(v any) error {
	return json.NewEncoder(r.w).Encode(v)
}

func (r *response) setHeaders(headers Headers) {
	r.w.Header().Set("Content-Type", "application/json")
	for k, v := range headers {
		r.w.Header().Set(k, v)
	}
}

func (r *response) Success(code int, message string, headers Headers, data any) error {
	r.setHeaders(headers)
	r.w.WriteHeader(code)
	v := success{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}
	return r.encode(v)
}

func (r *response) Fail(err error, headers Headers) error {
	code, message, stack := extractError(err)

	r.setHeaders(headers)
	r.w.WriteHeader(code)
	v := fail{
		StatusCode: code,
		Message:    message,
		Errors:     stack,
	}
	return r.encode(v)
}

func extractError(err error) (code int, message string, stack []string) {
	switch e := err.(type) {
	case *errorx.E:
		code = e.Code
		message = http.StatusText(code)
		stack = append(stack, e.Message)
	case validator.ValidationErrors:
		code = http.StatusBadRequest
		message = http.StatusText(code)
		for _, vErr := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("Error '%s' validation for '%s' field", vErr.Tag(), vErr.Field())
			stack = append(stack, errMsg)
		}
	default:
		code = http.StatusInternalServerError
		message = http.StatusText(code)
		stack = []string{e.Error()}
	}

	return
}

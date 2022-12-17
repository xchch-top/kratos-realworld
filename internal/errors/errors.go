package errors

import (
	"errors"
	"fmt"
)

type HttpError struct {
	Errors map[string]any `json:"errors"`
	Code   int            `json:"-"`
}

func NewHttpError(code int, field string, detail any) *HttpError {
	httpError := &HttpError{
		Code:   code,
		Errors: make(map[string]any),
	}
	httpError.Errors[field] = detail

	return httpError
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("FromError %d", e.Code)
}

func FromError(err error) *HttpError {
	if err == nil {
		return nil
	}
	if se := new(HttpError); errors.As(err, &se) {
		return se
	}
	return &HttpError{}
}

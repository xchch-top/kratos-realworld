package errors

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
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
	// 是HttpError类型
	if se := new(HttpError); errors.As(err, &se) {
		return se
	}
	// 是kratos封装的error
	if se := new(errors.Error); errors.As(err, &se) {
		return NewHttpError(int(se.Code), se.Reason, se.Message)
	}
	// 是原生error
	return NewHttpError(http.StatusInternalServerError, "internal", err.Error())
}

package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type (
	Error           string
	ValidationError []FieldError
	FieldError      struct {
		Field   string `json:"field,omitempty"`
		Error   string `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}
	AvailableErrors []ErrorResponse

	ResponseError struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}
)

func HandleErrorProvider(body io.ReadCloser) (ResponseError, error) {
	var res ResponseError
	err := json.NewDecoder(body).Decode(&res)
	if err != nil {
		return ResponseError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Error:   err,
		}, err
	}

	return res, nil
}

func (e Error) Error() string {
	return string(e)
}

var (
	ErrUnauthenticated      = Error("unauthenticated")
	ErrUnauthorized         = Error("unauthorized")
	ErrInvalidInterfaceType = Error("invalid type")
	ErrInvalidParameter     = Error("invalid parameter")
	ErrInvalidActionType    = Error("invalid action_type")
	ErrForbidden            = Error("you dont have access")
	ErrNotFound             = Error("not found")
	ErrUnprocessable        = Error("unproccesable")
	ErrBadRequest           = Error("bad request")
	ErrInternalServer       = Error("internal server error")
)

func DefaultAvailableErrors() AvailableErrors {
	return []ErrorResponse{
		{
			Code:    http.StatusNotFound,
			Err:     sql.ErrNoRows,
			Message: "data not found",
		},
		{
			Code:    http.StatusNotFound,
			Err:     ErrInvalidActionType,
			Message: "action type not found",
		},
		{
			Code:    http.StatusForbidden,
			Err:     ErrForbidden,
			Message: "you dont have access",
		},
		{
			Code:    http.StatusUnauthorized,
			Err:     ErrUnauthenticated,
			Message: "unauthenticated",
		},
		{
			Code:    http.StatusUnprocessableEntity,
			Err:     ErrInvalidInterfaceType,
			Message: "",
		},
		{
			Code:    http.StatusUnprocessableEntity,
			Err:     ErrInvalidParameter,
			Message: "",
		},
		{
			Code:    http.StatusBadRequest,
			Err:     io.EOF,
			Message: "please check your input",
		},
		{
			Code:    http.StatusNotFound,
			Err:     ErrNotFound,
			Message: "data not found",
		},
		{
			Code:    http.StatusUnprocessableEntity,
			Err:     ErrUnprocessable,
			Message: "unprocessable",
		},
	}
}

type ErrorResponse struct {
	Code    int
	Message string
	Err     error
	ValidationError
}

func (a *AvailableErrors) CustomeError(err AvailableErrors) *AvailableErrors {
	*a = append(*a, err...)
	return a
}

func (a AvailableErrors) GetError(err error) ErrorResponse {
	for _, e := range a {
		if errors.Is(err, e.Err) {
			if e.Message == "" {
				e.Message = err.Error()
			}
			return e
		}
	}
	causer := errors.Cause(err)

	switch causer {
	default:
		switch causer.(type) {
		case validation.Errors:
			return ErrorResponse{
				Code:    http.StatusUnprocessableEntity,
				Message: "invalid Request",
				Err:     err,
			}
		}
	}

	return ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Err:     err,
	}
}

func GetError(err error) ErrorResponse {
	for _, e := range DefaultAvailableErrors() {
		if errors.Is(err, e.Err) {
			if e.Message == "" {
				e.Message = err.Error()
			}
			return e
		}
	}
	causer := errors.Cause(err)
	switch causer {
	default:
		switch causer.(type) {
		case validation.Errors:
			return ErrorResponse{
				Code:    http.StatusUnprocessableEntity,
				Message: "invalid Request",
				Err:     err,
			}
		}
	}

	return ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Err:     err,
	}
}

// ErrInvalidMetaData for error invalid meta data
func ErrInvalidMetaData(additional string) Error {
	return Error(fmt.Sprintf("Invalid Meta Data %s", additional))
}

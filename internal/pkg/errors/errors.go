package errors

import (
	"errors"
	"net/http"
	"reflect"
	"uri-shortener/internal/pkg/errors/router_errors"
)

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type RequestError struct {
	Error ErrorInfo `json:"error"`
}

var mapErrorToCode = map[error]int{
	router_errors.BadRequest: http.StatusBadRequest,
	router_errors.NotFound:   http.StatusNotFound,
}

var mapErrorToCodeString = map[error]string{
	router_errors.BadRequest: "bad_request",
	router_errors.NotFound:   "not_found",
}

func getRootError(err error) error {
	currentErr := err
	for errors.Unwrap(currentErr) != nil {
		currentErr = errors.Unwrap(currentErr)
	}

	return currentErr
}

func resolveErrorToCode(err error) int {
	var code int
	isErrorFound := false
	rootErr := getRootError(err)

	if reflect.TypeOf(rootErr).Comparable() {
		code, isErrorFound = mapErrorToCode[rootErr]
	}

	if !isErrorFound {
		code = http.StatusInternalServerError
	}

	return code
}

func resolveErrorToCodeString(err error) string {
	code, ok := mapErrorToCodeString[err]
	if !ok {
		return "internal_error"
	}

	return code
}

func GetErrorMessageAndCode(err error) (*RequestError, int) {
	rootErr := getRootError(err)

	code := resolveErrorToCode(rootErr)
	codeString := resolveErrorToCodeString(rootErr)

	return &RequestError{
		Error: ErrorInfo{
			Code:    codeString,
			Message: err.Error(),
		},
	}, code
}

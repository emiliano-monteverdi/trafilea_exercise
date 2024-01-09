package errors

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Error interface {
	error
	StatusCode() int
	Code() string
	Message() string
	Cause() Cause
}

type _error struct {
	status  int
	code    string
	message string
	cause   Cause
}

type Cause struct {
	code        string
	description string
	error       error
}

func (e _error) MarshalJSON() ([]byte, error) {
	_cause := struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
		Error       error  `json:"error,omitempty"`
	}{
		Code:        e.cause.code,
		Description: e.cause.description,
		Error:       e.cause.error,
	}

	_err := struct {
		Status  int         `json:"status"`
		Code    string      `json:"error"`
		Message string      `json:"message"`
		Cause   interface{} `json:"cause"`
	}{
		Status:  e.status,
		Code:    e.code,
		Message: e.message,
		Cause:   _cause,
	}

	return json.Marshal(_err)
}

func (e _error) StatusCode() int {
	return e.status
}

func (e _error) Code() string {
	return e.code
}

func (e _error) Message() string {
	return e.message
}

func (e _error) Cause() Cause {
	return e.cause
}

func (e _error) Error() string {
	bytesArray, _ := e.MarshalJSON()
	return string(bytesArray)
}

func New(status int, message string, cause ...Cause) Error {
	_cause := Cause{}
	if cause != nil {
		_cause = cause[0]
	}
	return _error{
		status:  status,
		code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		message: message,
		cause:   _cause,
	}
}

func NewCauseError(err error) Cause {
	return Cause{error: err}
}

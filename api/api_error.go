package api

import (
	"encoding/json"
	"net/http"
)

func Error(code ModuleError) APIError {
	var err APIError
	err.Module = code.Module
	return err.WithCode(code.Error)
}

type APIError struct {
	status  int         `json:"-"`
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
	Module  string      `json:"module"`
}

func (e APIError) StatusCode() int {
	return e.status
}

func (e APIError) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e APIError) WithStatus(status int) APIError {
	e.status = status
	return e
}

func (e APIError) WithCode(code Code) APIError {
	e.Code = code
	e.Message = string(code)
	if status, ok := codeStatusMap[e.Code]; ok {
		e.status = status
	} else {
		e.status = http.StatusBadRequest
	}
	return e
}

func (e APIError) WithMessage(message string) APIError {
	e.Message = message
	return e
}

func (e APIError) WithDetails(details interface{}) APIError {
	e.Details = details
	return e
}

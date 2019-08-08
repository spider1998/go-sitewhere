package api

import (
	"google.golang.org/genproto/googleapis/rpc/code"
	"net/http"
)

type ModuleError struct {
	Module string
	Error  Code
}

type Code string

const (
	InternalServerError Code = "INTERNAL_SERVER_ERROR"
	NotFound            Code = "NOT_FOUND"
	MethodNotAllowed    Code = "METHOD_NOT_ALLOWED"

	InvalidData           Code = "INVALID_DATA"
	InvalidToken          Code = "INVALID_TOKEN"
	InvalidPersonToken    Code = "INVALID_PERSON_TOKEN"
	InvalidRolePermission Code = "INVALID_ROLE_PERMISSION"
	InvalidDeviceKey      Code = "INVALID_DEVICE_KEY"
)

var codeStatusMap = map[Code]int{
	InternalServerError: http.StatusInternalServerError,
	NotFound:            http.StatusNotFound,
	MethodNotAllowed:    http.StatusMethodNotAllowed,
}

//grpc service错误码
const (
	WeakPasswordCode         code.Code = 100
	OperatorExistCode        code.Code = 101
	OperatorNotExistCode     code.Code = 102
	InvalidPasswordCode      code.Code = 103
	InvalidOperatorTokenCode code.Code = 104
)

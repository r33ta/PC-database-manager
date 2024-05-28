package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"` // Error, Ok
	Error  string `json:"error,omitempty"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		case "name":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid name", err.Field()))
		case "ram_id":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid ram id", err.Field()))
		case "cpu_id":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid cpu id", err.Field()))
		case "gpu_id":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid gpu id", err.Field()))
		case "memory_id":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid memory id", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}

package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMssgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMssgs = append(errMssgs, fmt.Sprintf("Field %s is required", err.Field()))
		default:
			errMssgs = append(errMssgs, fmt.Sprintf("Field %s is invalid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMssgs, ","),
	}
}

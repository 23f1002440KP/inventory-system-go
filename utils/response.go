package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "ok"
	StatusError = "error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json") 
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response{
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response{
	var errors []string
	for _,err := range errs {
		switch err.ActualTag() {
		case "required":
			errors = append(errors, err.Field() + " is required")
		default :
			errors = append(errors, err.Field() + " is not valid")

				}		}
	return Response{
		Status: StatusError,
		Error: strings.Join(errors, ", "),
	}
}
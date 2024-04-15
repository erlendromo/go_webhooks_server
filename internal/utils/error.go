package utils

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	SystemError string `json:"systemError"`
	Statuscode  int    `json:"statuscode"`
}

func NewErrorResponse(err error, sc int) ErrorResponse {
	if err == nil {
		err = errors.New("internal server error")
	}

	if sc == 0 {
		sc = http.StatusInternalServerError
	}

	return ErrorResponse{
		SystemError: err.Error(),
		Statuscode:  sc,
	}
}

func ERROR(w http.ResponseWriter, err error, sc int) {
	JSON(w, sc, NewErrorResponse(err, sc))
}

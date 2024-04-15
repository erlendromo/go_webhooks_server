package utils

import (
	"encoding/json"
	"net/http"
	"webhooks/internal/constants"
)

func contentType(w http.ResponseWriter, contentType string, statuscode int) {
	w.Header().Set(constants.CONT_TYPE, contentType)
	WriteStatus(w, statuscode)
}

func WriteStatus(w http.ResponseWriter, statuscode int) {
	w.WriteHeader(statuscode)
}

func JSON(w http.ResponseWriter, statuscode int, data any) {
	contentType(w, constants.APP_JSON, statuscode)
	json.NewEncoder(w).Encode(data)
}

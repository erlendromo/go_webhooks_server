package handlers

import (
	"errors"
	"html/template"
	"net/http"
	"webhooks/internal/constants"
	"webhooks/internal/utils"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmp, err := template.ParseFiles(constants.ROOT_HTML_PATH)
		if err != nil {
			utils.ERROR(w, err, http.StatusInternalServerError)
			return
		}

		err = tmp.Execute(w, nil)
		if err != nil {
			utils.ERROR(w, err, http.StatusInternalServerError)
			return
		}
	default:
		utils.ERROR(w, errors.New(constants.METHOD_NOT_ALLOWED), http.StatusMethodNotAllowed)
	}
}

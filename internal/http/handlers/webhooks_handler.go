package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
	"webhooks/internal/business/domains"
	"webhooks/internal/constants"
)

func WebhooksHandler(w http.ResponseWriter, r *http.Request, wUC domains.WebhookUsecase) {
	switch r.Method {
	case http.MethodGet:
		displayHTML(w, wUC)
	case http.MethodPost:
		addWebhook(w, r, wUC)
	default:
		http.Error(w, constants.METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}

// TODO Prune documents that have been stored for more than (1 hour?)
func displayHTML(w http.ResponseWriter, wUC domains.WebhookUsecase) {

	// Move this to business domains-usecases???
	tmp, err := template.ParseFiles(constants.INDEX_HTML_PATH)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	whs, s, err := wUC.Get()
	if err != nil {
		http.Error(w, err.Error(), s)
		return
	}

	webhooks := map[string][]domains.Webhook{
		"Webhooks": whs,
	}

	err = tmp.Execute(w, webhooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addWebhook(w http.ResponseWriter, r *http.Request, wUC domains.WebhookUsecase) {
	// Extract this into own method??
	var v any
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO This implementation will be reworked when setup on other service is ready
	wh := domains.Webhook{
		Method:    r.Method,
		Url:       fmt.Sprintf("%s%s", constants.OTHER_SERVICE, r.URL.Path),
		Timestamp: time.Now().Local().Format(constants.TIME_FORMAT),
		Content:   v,
	}

	s, err := wUC.Store(&wh)
	if err != nil {
		http.Error(w, err.Error(), s)
	}
}

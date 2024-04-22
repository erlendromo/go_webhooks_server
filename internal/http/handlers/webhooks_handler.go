package handlers

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"webhooks/internal/business/domains"
	"webhooks/internal/business/usecases"
	"webhooks/internal/constants"
	"webhooks/internal/utils"

	"cloud.google.com/go/firestore"
)

func WebhooksHandler(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	wUC := usecases.NewDBWebhook(client)

	switch r.Method {
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		displayHTML(w, r, wUC)
	case http.MethodPost:
		addWebhook(w, r, wUC)
	default:
		utils.ERROR(w, errors.New(constants.METHOD_NOT_ALLOWED), http.StatusMethodNotAllowed)
	}
}

// TODO Prune documents that have been stored for more than (1 hour?)
func displayHTML(w http.ResponseWriter, r *http.Request, wUC domains.WebhookUsecase) {
	tmp, err := template.ParseFiles(constants.INDEX_HTML_PATH)
	if err != nil {
		utils.ERROR(w, err, http.StatusInternalServerError)
		return
	}

	whs, sc, err := wUC.Get(r.Context())
	if err != nil {
		utils.ERROR(w, err, sc)
		return
	}

	webhooks := map[string][]domains.Webhook{
		"Webhooks": whs,
	}

	err = tmp.Execute(w, webhooks)
	if err != nil {
		utils.ERROR(w, err, http.StatusInternalServerError)
		return
	}
}

func addWebhook(w http.ResponseWriter, r *http.Request, wUC domains.WebhookUsecase) {
	var n struct {
		ID      string `json:"id"`
		Country string `json:"country"`
		Event   string `json:"event"`
		Time    string `json:"time"`
	}
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		utils.ERROR(w, err, http.StatusInternalServerError)
		return
	}

	wh := domains.Webhook{
		Event:     n.Event,
		Url:       r.RemoteAddr,
		Timestamp: n.Time,
	}

	sc, err := wUC.Store(r.Context(), &wh)
	if err != nil {
		utils.ERROR(w, err, sc)
		return
	}

	utils.JSON(w, sc, "Webhook-trigger recieved!")
}

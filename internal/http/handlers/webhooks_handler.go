package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
	"webhooks/internal/business/domains"
	"webhooks/internal/business/usecases"
	"webhooks/internal/constants"

	"cloud.google.com/go/firestore"
)

// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /accounts [get]
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
		http.Error(w, constants.METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
	}
}

// TODO Prune documents that have been stored for more than (1 hour?)
func displayHTML(w http.ResponseWriter, r *http.Request, wUC domains.WebhookUsecase) {

	// Move this to business domains-usecases???
	tmp, err := template.ParseFiles(constants.INDEX_HTML_PATH)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	whs, s, err := wUC.Get(r.Context())
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

	s, err := wUC.Store(r.Context(), &wh)
	if err != nil {
		http.Error(w, err.Error(), s)
	}
}

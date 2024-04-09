package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"
	"webhooks/internal/business/domains"
	"webhooks/internal/datasources/database"

	"cloud.google.com/go/firestore"
)

func HandleHTML(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	switch r.Method {
	case http.MethodGet:
		displayHTML(w, r, client)
	case http.MethodPost:
		addWebhook(w, r, client)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func displayHTML(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	tmp, err := template.ParseFiles("internal/html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := database.GetDocuments[domains.Webhook](client, r.Context(), "webhook_triggers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	webhooks := map[string][]domains.Webhook{
		"Webhooks": resp,
	}

	// This is how the webhooks will be utilized by the application
	// wUC.Get()
	// wUC.GetByID("id...")
	// wUC.Store(nil)

	err = tmp.Execute(w, webhooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addWebhook(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	timeFormat := "2006-01-02 15:04"

	var v any
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	wh := domains.Webhook{
		Method:    r.Method,
		Url:       r.URL.Path,
		Timestamp: time.Now().Local().Format(timeFormat),
		Content:   v,
	}

	err = database.UploadDocument(client, r.Context(), "webhook_triggers", wh)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

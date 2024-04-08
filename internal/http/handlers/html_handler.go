package handlers

import (
	"html/template"
	"net/http"
	"webhooks/internal/business/domains"
)

func HandleHTML(w http.ResponseWriter, r *http.Request, wUC domains.WebhookUsecase) {

	tmp, err := template.ParseFiles("internal/html/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// This is temporary - will all be dynamically retrieved from the other service
	webhooks := map[string][]domains.Webhook{
		"Webhooks": {
			{ID: "1", Url: "https://localhost:8080/dashboard/v1/registrations/", Country: "", Event: "REGISTER"},
			{ID: "2", Url: "https://localhost:8080/dashboard/v1/dashboards/", Country: "", Event: "REGISTER"},
		},
	}

	// This is how the webhooks will be utilized by the application
	// wUC.Get()
	// wUC.GetByID("id...")
	// wUC.Store(nil)

	err = tmp.Execute(w, webhooks)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

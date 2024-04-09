package routes

import (
	"encoding/json"
	"net/http"
	"webhooks/internal/config"
	"webhooks/internal/constants"
	"webhooks/internal/http/handlers"

	"cloud.google.com/go/firestore"
)

type Router struct {
	Port   string
	Client *firestore.Client
}

func NewRouter(config config.Config, client *firestore.Client) *Router {
	return &Router{
		Port:   config.Port,
		Client: client,
	}
}

// Should Handlerfunctions recieve the client as well or a middleware?
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//wUC := usecases.NewWebhookUsecase(router.Client, r.Context())

	binder := BindRequest(r)

	switch binder.Endpoint {
	case constants.ROOT:
		json.NewEncoder(w).Encode("Root")
	case constants.WEBHOOKS_PATH:
		handlers.HandleHTML(w, r, router.Client)
	default:
		http.Error(w, "invalid endpoint", http.StatusNotFound)
	}
}

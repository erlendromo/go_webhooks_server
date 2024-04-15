package routes

import (
	"net/http"
	"webhooks/internal/config"
	"webhooks/internal/constants"
	"webhooks/internal/http/handlers"

	"cloud.google.com/go/firestore"
	swaggerFiles "github.com/swaggo/files"
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

// TODO Should Handlerfunctions recieve the usecase, or a middleware?
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	binder := BindRequest(r)

	// TODO Add default page on root path??
	switch binder.Endpoint {
	case constants.WEBHOOKS_PATH:
		handlers.WebhooksHandler(w, r, router.Client)
	case constants.SWAGGER_PATH:
		swaggerFiles.Handler.ServeHTTP(w, r)
	default:
		http.Error(w, constants.INVALID_ENDPOINT, http.StatusNotFound)
	}
}

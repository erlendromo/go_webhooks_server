package routes

import (
	"encoding/json"
	"net/http"
	"webhooks/internal/config"
	"webhooks/internal/constants"
)

type Router struct {
	Port string
}

func NewRouter(config config.Config) *Router {
	return &Router{
		Port: config.Port,
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	binder := BindRequest(w, r)
	switch binder.Endpoint {
	case constants.WEBHOOKS_PATH:
		json.NewDecoder(r.Body).Decode(&binder)
	default:
		http.Error(w, "invalid endpoint", http.StatusNotFound)
	}
}

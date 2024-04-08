package server

import (
	"fmt"
	"log"
	"net/http"
	"webhooks/internal/config"
	"webhooks/internal/http/routes"
)

func StartServer() {
	config := config.NewConfig()
	router := routes.NewRouter(*config)

	log.Printf("Server started on port %s...\n", router.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", router.Port), router)
}

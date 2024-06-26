package server

import (
	"fmt"
	"log"
	"net/http"
	"webhooks/internal/config"
	"webhooks/internal/constants"
	"webhooks/internal/datasources/database"
	"webhooks/internal/http/routes"
)

func StartServer() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	client, err := database.NewFirestoreClient(*config)
	if err != nil {
		log.Fatal(constants.CONNECTION_ERR)
	}

	router := routes.NewRouter(*config, client)

	log.Printf("Server started on port %s...\n", router.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", router.Port), router))
}

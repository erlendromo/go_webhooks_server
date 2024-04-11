package server

import (
	"fmt"
	"log"
	"net/http"
	"webhooks/internal/config"
	"webhooks/internal/datasources/firestoredb"
	"webhooks/internal/http/routes"
)

func StartServer() {
	config := config.NewConfig()

	client, err := firestoredb.NewFirestoreClient(*config)
	if err != nil {
		log.Fatal("Connection to firebase was unsuccessful")
	}

	router := routes.NewRouter(*config, client)

	log.Printf("Server started on port %s...\n", router.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", router.Port), router))
}

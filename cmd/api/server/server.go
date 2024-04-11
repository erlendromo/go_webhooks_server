package server

import (
	"fmt"
	"log"
	"net/http"
	"webhooks/internal/config"
	"webhooks/internal/datasources/firestoredb"
	"webhooks/internal/http/routes"

	"github.com/joho/godotenv"
)

func StartServer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	client, err := firestoredb.NewFirestoreClient(*config)
	if err != nil {
		log.Fatal("Connection to firebase was unsuccessful")
	}

	router := routes.NewRouter(*config, client)

	log.Printf("Server started on port %s...\n", router.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", router.Port), router))
}

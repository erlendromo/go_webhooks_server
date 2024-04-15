package main

import (
	"webhooks/cmd/server"

	"webhooks/docs"
)

func init() {
	docs.SwaggerInfo.Title = "dhgjslk"
	docs.SwaggerInfo.Description = "hgdksj"
	docs.SwaggerInfo.Version = "cjdhbkfhj"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}

func main() {
	server.StartServer()
}

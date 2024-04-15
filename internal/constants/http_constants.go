package constants

// Change this if used by other url
const OTHER_SERVICE = "http://localhost:8080"

const (
	ROOT     = "/"
	WEBHOOKS = "webhooks"
	SWAGGER  = "swagger"
)

const (
	WEBHOOKS_PATH = ROOT + WEBHOOKS + ROOT
	SWAGGER_PATH  = ROOT + SWAGGER + ROOT
)

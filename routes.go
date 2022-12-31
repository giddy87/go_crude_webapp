package main

import (
	"net/http"
	"pkg/config"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := Pat.New()

}

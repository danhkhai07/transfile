package app

import (
	"net/http"

	"transfile/config"
)

func addRoutes(
	mux *http.ServeMux,
	cfg *config.Config,
	logger Logger,
) {
	mux.HandleFunc("GET /health", 			getHealth())
	mux.HandleFunc("GET /lookup/{hash}", 	lookupFile())
	mux.HandleFunc("GET /download/{hash}", 	downloadFile())
	mux.HandleFunc("POST /upload", 			postFile())
}

package server

import (
	"net/http"

	"transfile/config"
)

func addRoutes(
	svr *Server,
	mux *http.ServeMux,
	cfg *config.Config,
	logger Logger,
) {
	mux.HandleFunc("GET /health", 			svr.getHealth)
	mux.HandleFunc("GET /lookup/{hash}", 	svr.lookupFile)
	mux.HandleFunc("GET /download/{hash}", 	svr.downloadFile)
	mux.HandleFunc("POST /upload", 			svr.postFile)
}

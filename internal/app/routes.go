package app

import (
	"net/http"

	"transfile/config"
)

func addRoutes(
	mux *http.ServeMux,
	cfg *config.Config,
	logger config.Logger,
) {
	mux.Handle("/", mockHome())
}

func mockHome() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello world!"))
		},
	)
}

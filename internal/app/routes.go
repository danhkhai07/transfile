package app

import (
	"net/http"

	"github.com/danhkhai07/transfile/config"
)

func addRoutes(
	mux *http.ServeMux,
	cfg *config.Config,
	logger config.Logger,
) {
	mux.Handle("/", )
}

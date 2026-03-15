package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"transfile/config"
)

type Server struct {
	addr string
	logger Logger
	httpServer *http.Server
}

func (server *Server) Run(
	ctx context.Context,
	args []string,
	getenv func(string) string,
) (err error) {
	ctx, osCancel := signal.NotifyContext(ctx, os.Interrupt)
	defer osCancel()

	// errChan := make(chan error, 1)
	/*
		check flags & get environment variables here
	*/


	go func() {
		server.logger.Writeln("listening on %s", server.addr)
		err := server.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			server.logger.Errwriteln("error listen and serve: %s", err)
			return
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, shutdownCancel := context.WithTimeout(shutdownCtx, 10 * time.Second)
		defer shutdownCancel()

		server.logger.Writeln("server shutting down...")
		if err := server.httpServer.Shutdown(shutdownCtx); err != nil {
			server.logger.Errwriteln("error shutting down http server: %s", err)
			return
		}
	}()
	wg.Wait()
	
	return nil
}

func NewServer(
	cfg *config.Config,
	logger Logger,
) (*Server) {
	server := Server{
		addr: net.JoinHostPort("0.0.0.0", cfg.Port),
		logger: logger,
		httpServer: &http.Server{},
	}
	server.httpServer.Addr = server.addr

	handler := NewHandler(
		&server,
		cfg,
		logger,
	)
	server.httpServer.Handler = handler

	return &server
}

func NewHandler(
	svr *Server,
	cfg *config.Config,
	logger Logger,
) (http.Handler) {
	mux := http.NewServeMux()
	addRoutes(
		svr,
		mux,
		cfg,	
		logger,
	)
	handler := http.Handler(mux)
	// add middleware

	return handler
}


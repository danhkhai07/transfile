package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"transfile/config"
	"transfile/internal/cache"
)

var (
	ErrInternalError = errors.New("internal error")
)

type Server struct {
	addr string
	logger Logger
	fileStore *cache.FileStore
	httpServer *http.Server
	aliveAt time.Time
}

func (svr *Server) Run(
	ctx context.Context,
	args []string,
	getenv func(string) string,
) {
	ctx, osCancel := signal.NotifyContext(ctx, os.Interrupt)
	defer osCancel()

	// errChan := make(chan error, 1)
	/*
		check flags & get environment variables here
	*/


	go func() {
		svr.logger.Writeln("listening on %s", svr.addr)
		svr.aliveAt = time.Now()
		err := svr.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			svr.logger.Errwriteln("error listen and serve: %s", err)
			log.Fatal(err)
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

		svr.logger.Writeln("server shutting down...")
		if err := svr.httpServer.Shutdown(shutdownCtx); err != nil {
			svr.logger.Errwriteln("error shutting down http server: %s", err)
			log.Fatal(err)
			return
		}
	}()
	wg.Wait()
}

func NewServer(
	cfg *config.Config,
	logger Logger,

) (*Server) {
	svr := Server{
		addr: net.JoinHostPort("0.0.0.0", cfg.Port),
		logger: logger,
		fileStore: cache.NewFileStore(),
		httpServer: &http.Server{},
	}
	svr.httpServer.Addr = svr.addr

	handler := NewHandler(
		&svr,
		cfg,
		logger,
	)
	svr.httpServer.Handler = handler

	return &svr
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

func (svr *Server) Uptime() int {
	return int(time.Now().Unix()) - int(svr.aliveAt.Unix())
}

func (svr *Server) Encode(w http.ResponseWriter, r *http.Request, status int, v any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
    }
    return nil
}

func (svr *Server) Decode(r *http.Request, v any) (error) {
    if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
    }
    return nil
}

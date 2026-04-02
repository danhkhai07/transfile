package main

import (
	"context"
	"os"

	"transfile/config"
	"transfile/internal/app/server"
)

func main() {
	cfg := config.ServerConfig {
		Port: "8080",
	}
	logger := config.StdLogger{}

	ctx := context.Background()
	server := server.NewServer(
		&cfg,
		&logger,
	)
	server.Run(ctx, os.Args, os.Getenv)
}

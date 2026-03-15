package main

import (
	"context"
	"os"

	"transfile/config"
	"transfile/internal/api"
)

func main() {
	cfg := config.Config{
		Port: "8080",
	}
	logger := config.StdLogger{}

	ctx := context.Background()
	server := app.NewServer(
		&cfg,
		&logger,
	)
	if err := server.Run(ctx, os.Args, os.Getenv); err != nil {
		os.Exit(1)
	}
}

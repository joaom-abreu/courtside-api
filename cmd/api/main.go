package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joaom-abreu/courtside-api/internal/config"
	"github.com/joaom-abreu/courtside-api/internal/database"
	"github.com/joaom-abreu/courtside-api/internal/handler"
	"github.com/joaom-abreu/courtside-api/internal/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := run(logger); err != nil {
		logger.Error("api stopped with error", "error", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	logger.Info("connected to postgres")

	router := handler.NewRouter(handler.NewHealthHandler(db))
	srv := server.New(cfg.HTTPAddr(), router, logger)

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- srv.Start()
	}()

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	return srv.Shutdown(shutdownCtx)
}

package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	app, err := NewApp(os.Getenv)
	if err != nil {
		slog.Error(fmt.Sprintf("init app:", err))
		os.Exit(1)
	}

	err = app.Start()
	if err != nil {
		slog.Error(fmt.Sprintf("starting app: %w", err))
		os.Exit(1)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, os.Kill)
	<-sigch
	slog.Info("recieved signal, shutting down")

	err = app.Shutdown()
	if err != nil {
		slog.Error(fmt.Sprintf("shutdown: %w", err))
	}
}

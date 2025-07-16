package main

import (
	"log/slog"
	"os"

	"github.com/ueno-tecnologia-org/go-core/pkg/logging"
)

func initLogger() {
	logConfig := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logJsonHandler := slog.NewJSONHandler(os.Stdout, &logConfig)

	customHandler := logging.NewCustomHandler(logJsonHandler)

	slog.SetDefault(slog.New(customHandler))
}

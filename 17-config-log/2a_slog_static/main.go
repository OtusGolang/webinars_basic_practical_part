package main

import (
	"errors"
	"log/slog"
	"os"
)

func main() {
	demoDefaultLogger()
}

func demoDefaultLogger() {
	// default logger
	//ctx := context.Background()
	slog.Debug("debug 1", "count", 3)
	slog.Info("info 1", slog.Int("count", 3), "hi", "there")
	slog.Error("oh oh", "error", errors.New("something bad happen"))

	// can change default logger
	logConfig := &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	logHandler := slog.NewJSONHandler(os.Stderr, logConfig)
	//logHandler := slog.NewTextHandler(os.Stderr, logConfig)

	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	slog.Debug("debug 2", "count", 3)
	slog.Info("info 2", "count", 3)
}

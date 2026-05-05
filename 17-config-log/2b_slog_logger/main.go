package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

func main() {
	demoLoggers()
}

func demoLoggers() {
	// setup:
	file, err := os.Create("main.log")
	if err != nil {
		panic(fmt.Errorf("create log file: %w", err))
	}
	defer file.Close()

	writer := io.MultiWriter(file, os.Stderr)
	logger := slog.New(slog.NewTextHandler(writer, nil))

	logger.Info("starting processing")

	// examples:
	requestLogger := logger.With("request_id", "12345")
	requestLogger.Info("going to call handler")
	err = handleRequest(requestLogger)
	//err = processRequest(requestLogger.WithGroup("handler")) // grouping is possible
	if err != nil {
		logger.Error("error while process request", "error", err)
	}
	logger.Info("processing finished")
}

func handleRequest(logger *slog.Logger) error {
	logger.Info("Processing started")
	logger.Warn("Processing took too long", slog.Duration("duration", 12*time.Second))
	return fmt.Errorf("some error")
}

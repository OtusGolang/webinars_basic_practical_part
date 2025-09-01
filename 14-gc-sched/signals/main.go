package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// print pid
	println("my PID:", os.Getpid())

	// Create a channel to receive signals
	signals := make(chan os.Signal, 1)
	// Notify the channel of the signals to listen for
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGHUP)
	//signal.Ignore(syscall.SIGHUP)
	for s := range signals {
		// Print the signal received
		fmt.Printf("Got signal: %v\n", s)
		if s == syscall.SIGINT {
			println("SIGINT received, exiting")
			break
		}
	}

}

// example of kill commands:
// kill -SIGINT <pid>
// kill -SIGTERM <pid>
// kill -SIGKILL <pid>

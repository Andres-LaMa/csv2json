package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// HandleShutdown ловит SIGINT/SIGTERM и выполняет callback.
func HandleShutdown(callback func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		callback()
		os.Exit(0)
	}()
}

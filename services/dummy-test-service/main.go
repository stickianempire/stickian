package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// The context should be the one controlling the lifecycle of the program
	// ensure external SIGINT and SIGTERM are gracefully handled.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	serverHandler := serverHandler{
		db: &mockDB{},
	}

	serverHandler.init_db()

	server := &http.Server{
		Addr:              ":4000",
		ReadHeaderTimeout: 3 * time.Second,
	}
	http.HandleFunc("/", serverHandler.handleRequest)

	go func() {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("server closed")
		} else if err != nil {
			log.Printf("error starting server: %v", err)
		}
	}()
	<-ctx.Done()
}

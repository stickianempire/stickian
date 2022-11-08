package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	// The context should be the one controlling the lifecycle of the program
	// ensure external SIGINT and SIGTERM are gracefully handled.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := serverHandler{
		db: &mockDB{},
	}

	server.init_db()

	// We're showing that we can connect to MongoDB
	func() {
		client, ctx, cancel, err := connect(ctx, "mongodb://localhost:27017")
		if err != nil {
			log.Panicf("could not connect %v", err)
		}
		defer close(client, ctx, cancel)

		ping(client, ctx)
	}()

	http.HandleFunc("/", server.handleRequest)

	err := http.ListenAndServe(":4000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed")
	} else if err != nil {
		log.Printf("error starting server: %v", err)
	}
}

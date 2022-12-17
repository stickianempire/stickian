package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	gamecityview "github.com/stickianempire/stickian/services/game-city-view/service"
)

func main() {
	// The context should be the one controlling the lifecycle of the program
	// ensure external SIGINT and SIGTERM are gracefully handled.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	dbClient, err := gamecityview.NewDatabaseClient(ctx, 5432, "localhost", "postgres", "postgres")
	if err != nil {
		// nolint // we really want to die if db connection fails
		log.Fatalf("db client creation %v", err)
	}
	defer dbClient.Close()

	serverHandler := gamecityview.NewServerHandler(ctx, dbClient)

	server := &http.Server{
		Addr:              ":4000",
		ReadHeaderTimeout: 3 * time.Second,
	}
	http.HandleFunc("/city", serverHandler.HandleRequest)

	go func() {
		err := server.ListenAndServe()
		switch {
		case errors.Is(err, http.ErrServerClosed):
			log.Printf("server closed")
		default:
			log.Printf("left ListenAndServe with err: %v", err)
		}
	}()

	// wait for the shutdown signal from the outside
	<-ctx.Done()

	shutdownCtx, shutdownStop := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownStop()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("shuting down got %v", err)
	}
}

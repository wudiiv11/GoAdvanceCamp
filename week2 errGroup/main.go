package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})

	stop := make(chan struct{})
	mux.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
		stop <- struct{}{}
	})

	server := http.Server{Handler: mux, Addr: "127.0.0.1:8080"}
	group.Go(func() error {
		return server.ListenAndServe()
	})

	group.Go(func() error {
		timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Println("ctx is done")
		case <-stop:
			log.Println("server will shutdown")
		case <-quit:
			log.Println("process killed")
		}

		return server.Shutdown(timeoutCtx)

	})

	if err := group.Wait(); err != nil {
		log.Fatal(err)
	}
}

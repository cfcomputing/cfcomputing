package main

import (
	"cfcomputing/api"
	"cfcomputing/auth"
	"cfcomputing/router"
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func initServer() *http.Server {
	port := flag.String("port", ":80", "http service address")
	server := &http.Server{Addr: *port}

	http.HandlerFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to cfcomputing.com! Check back soon there is work being done\n")
	})

	// going to add some connections to redis and postgresql etc into the api
	api := new(api.API)
	http.Handle("/api", router.ApplyMiddleware(api, auth.CheckAuth("api-general")))

	return server
}

func main() {
	srv := initServer()

	// taken from https://golang.org/pkg/net/http/#Server.Shutdown && https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		s := <-sigint

		log.Print("Shutting down")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}

		log.Printf("Shutdown complete %v", s)
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	s := <-idleConnsClosed
	log.Printf("idleconnsclosed %v", s)
}

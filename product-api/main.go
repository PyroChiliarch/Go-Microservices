package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"productapiHandlers"
	"time"
)

func main() {

	//Create handlers
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := productapiHandlers.NewProducts(l)

	//Register handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	//Start the web server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//Register to receive interrupt signals
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	//Block until interrupt
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}

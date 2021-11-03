package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abesheknarayan/microservice-go/handlers"
)

func main() {
	l := log.New(os.Stdout, "microservice-go", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)
	servermux := http.NewServeMux()
	servermux.Handle("/", helloHandler)
	servermux.Handle("/bye", goodbyeHandler)

	httpServer := &http.Server{
		Addr:         ":9000",
		Handler:      servermux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// anonymous goroutine
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// signal channel notifies OS signals like kill / interrupt
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	signal := <-signalChannel
	l.Println("Gracefully Terminating Server, got signal", signal)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// gracefully shutting down the http Server
	httpServer.Shutdown(timeoutContext)
}

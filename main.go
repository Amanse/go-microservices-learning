package main

import (
	"context"
	"github.com/Amanse/server/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	//Import a new logger to pass
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//post handler requires the logger, pass it in
	ph := handlers.NewProduct(l)

	//Create a new server mux
	sm := http.NewServeMux()
	//handle / of server to product handler
	sm.Handle("/", ph)

	//Add server parameters and also time out
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//Start server and show if any error
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Server running on port 9090")
	}()

	//Make signal chan
	sigChan := make(chan os.Signal)
	//Notify on Interrupt or kill
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	//if recieved signal gracefully shutdown and show which signal
	sig := <-sigChan
	l.Println("Recieved msg. graceful shudown", sig)

	//Shutdown after timeout after everything is done
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

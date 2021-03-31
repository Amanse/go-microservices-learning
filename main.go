package main

import (
	"context"
	"github.com/Amanse/server/handlers"
	"github.com/gorilla/mux"
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

	log.Println("Server running on port 9090")
	//Create a new server mux
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddleWareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddleWareProductValidation)

	// ------- MY CODE HERE BEWARE --------

	// Add delete Router and delete route here

	// ------- MY CODE ENDS ---------------
	//handle / of server to product handler
	//sm.Handle("/products", ph)

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

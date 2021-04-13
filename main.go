package main

import (
	"alert/crypto/markets_usercase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/markets/{ids}", markets_usercase.GetMarkets)
	myRouter.HandleFunc("/markets", markets_usercase.GetMarkets)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	port := ":8080"
	log.Println("Application Starting with Port "+port)
	log.Fatal(http.ListenAndServe(port, myRouter))
}


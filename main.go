package main

import (
	"alert/crypto/usercase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/markets/{ids}", usercase.GetMarkets)
	myRouter.HandleFunc("/markets", usercase.GetMarkets)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	port := ":8080"
	usercase.CronToTriggerWebhook()
	log.Println("Application Starting with Port " + port)
	log.Fatal(http.ListenAndServe(port, myRouter))
}

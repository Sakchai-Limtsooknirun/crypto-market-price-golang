package main

import (
	"alert/crypto/markets_usercase"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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

	log.Fatal(http.ListenAndServe(getPort(), myRouter))
}


func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku " + port)
	}
	return ":" + port
}

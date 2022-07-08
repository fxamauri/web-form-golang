package main

import (
	"log"
	"net/http"

	"github.com/fxamauri/web-form-golang/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)

	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
}

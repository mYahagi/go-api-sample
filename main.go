package main

import (
	"log"
	"net/http"

	"github.com/mYahagi/go-api-sample/app/controller"
)

func main() {
	http.HandleFunc("/price", controller.NewPriceController().Show)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Printf("process interrupted: %s \n", err)
		panic("an error occurred")
	}
}

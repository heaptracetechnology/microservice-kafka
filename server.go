package main

import (
	route "github.com/heaptracetechnology/microservice-kafka/route"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}

package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) string {
	return "123"
}

//The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
//ListenAndServe listens on TCP network on the provided port and calls the handler to handle requests on incoming connections
func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}

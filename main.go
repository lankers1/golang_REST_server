package main

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Person struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Score int16              `json:"score,omitempty" bson:"score,omitempty"`
}

var client *mongo.Client

//The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
//ListenAndServe listens on TCP network on the provided port and calls the handler to handle requests on incoming connections
func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}

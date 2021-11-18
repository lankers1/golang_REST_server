package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Score int                `json:"score,omitempty" bson:"score,omitempty"`
}

var client *mongo.Client

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Print("POST started")
	response.Header().Add("content-type", "application/json")
	person := Person{}
	_ = json.NewDecoder(request.Body).Decode(&person)

	collection := client.Database("testdatabase").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, person)
	fmt.Println(result, "result")

	if err != nil {
		panic("Error")
	}

	json.NewEncoder(response).Encode(result)
}

func main() {
	fmt.Println("Starting application...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI("mongodb://mongodb")
	client, _ = mongo.Connect(ctx, clientOpts)

	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	fmt.Println("server listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

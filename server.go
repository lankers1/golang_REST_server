package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) string
}

type PlayerServer struct {
	store PlayerStore
}

//ResponseWriter implements the IO Writer with the Write function.
//Request represents an HTTP request received by a server or to be sent by a client.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	//TrimPrefix removes "/players/" from URL path string to return the players name
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)

	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

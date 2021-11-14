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
	//TrimPrefix removes "/players/" from URL path string to return the players name

	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

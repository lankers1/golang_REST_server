package main

import (
	"fmt"
	"net/http"
	"strings"
)

//ResponseWriter implements the IO Writer with the Write function.
//Request represents an HTTP request received by a server or to be sent by a client.
func PlayerServer(w http.ResponseWriter, r *http.Request) {

	//TrimPrefix removes "/players/" from URL path string to return the players name
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(2, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}

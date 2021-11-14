package main

import (
	"fmt"
	"net/http"
)

//ResponseWriter implements the IO Writer with the Write function.
//Request represents an HTTP request received by a server or to be sent by a client.
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

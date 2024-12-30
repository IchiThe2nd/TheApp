package main

import (
	"log"
	"net/http"
)

// logs
func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}

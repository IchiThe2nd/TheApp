package main

import (
	"log"
	"net/http"
)

// logs
func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}

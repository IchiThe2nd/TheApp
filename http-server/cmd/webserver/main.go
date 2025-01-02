package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IchiThe2nd/TheApp/http-server"
	// "github.com/IchiThe2nd/TheApp/http-server/poker"
)

const dbFileName = "game.db.json"

// logs
func main() {

	store,close, err := poker.FileSystempPlayerStoreFromFile(dbFileName)
		 err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not list enm on portt 5000 %v", err)
	}
}
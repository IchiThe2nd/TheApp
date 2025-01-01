package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IchiThe2nd/TheApp/http-server/poker"
	// "github.com/IchiThe2nd/TheApp/http-server/poker"
)

const dbFileName = "game.db.json"

// logs
func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating filesystem palyer stopre , %v", err)
	}
	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not list enm on portt 5000 %v", err)
	}
}

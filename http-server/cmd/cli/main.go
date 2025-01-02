package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/IchiThe2nd/TheApp/http-server"
)

const dbFileName = "game.db.json"

func main() {

	store, close, err := poker.FileSystempPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Lets play Poker")
	fmt.Println("Type {name} wins to record a win")
	poker.NewCLI(store, os.Stdin, poker.BlindAlerterFunc(poker.StdOutAlerter)).PlayPoker()

}

package main

import (
	"fmt"
	"net/http"

	"github.com/CayenneLow/Codenames/src/config"
	"github.com/CayenneLow/Codenames/src/game"
	"github.com/CayenneLow/Codenames/src/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.Init()
	logger.Init()

	http.HandleFunc("/newgame", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Query())
		game.NewGame()
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

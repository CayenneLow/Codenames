package main

import (
	"net/http"

	"github.com/CayenneLow/Codenames/internal/config"
	"github.com/CayenneLow/Codenames/internal/game"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Init()
	games := make(map[uuid.UUID]game.GameState)

	http.HandleFunc("/newgame", func(w http.ResponseWriter, r *http.Request) {
		log.Debug(r.URL.Query())
		newGame := game.NewGame(cfg)
		games[newGame.GameID] = newGame
	})

	log.Info("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

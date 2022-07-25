package main

import (
	"fmt"
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

	log.Infof("Starting server at port %s", cfg.ServerPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ServerPort), nil); err != nil {
		log.Fatal(err)
	}
}

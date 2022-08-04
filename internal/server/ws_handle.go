package server

import (
	"context"
	"time"

	"github.com/CayenneLow/Codenames/internal/database"
	"github.com/CayenneLow/codenames-eventrouter/pkg/event"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func HandleWsMessage(ctx context.Context, message []byte, ws *websocket.Conn, db database.Database) {
	// TODO: Handle ctx Done
	e, err := event.FromJSON(message)
	if err != nil {
		log.Error("Error converting JSON to Event", log.Fields{
			"error": err,
		})
	}

	switch e.Type {
	case "joinGame":
		// Error thrown when gameid does not exist
		handleJoinGame(db, e)

	}
}

func handleJoinGame(db database.Database, e event.Event) {
	dbCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := db.ReadBoardByGameID(dbCtx, e.GameID)

	if err != nil {
		log.Error("Error ReadBoardByGameID", log.Fields{
			"event": e,
			"error": err,
		})
	}
}

package server

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CayenneLow/Codenames/internal/config"
	"github.com/CayenneLow/Codenames/internal/database"
	"github.com/CayenneLow/Codenames/internal/game"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func Start(cfg config.Config, db database.Database) {
	// Create session ID
	sessionID := uuid.NewString()

	// Init connection to EventRouter
	ws, err := connectToEventRouter(cfg)
	if err != nil {
		log.Error("Error connecting to event router", log.Fields{
			"error": err,
		})
	}

	// Read websocket loop
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		defer ws.Close()
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Error("Error occured reading from Websocket, shutting down websocket", log.Fields{
					"error": err,
				})
				return
			}
			log.Debug("Received message from EventRouter WS", log.Fields{
				"msg": message,
			})
		}
	}(ctx)

	// Exit on SIGTERM and Interrupt
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			cancel()
		case syscall.SIGTERM:
			cancel()
		}
	}()

	// REST API Server for handling new game
	http.HandleFunc("/newgame", func(w http.ResponseWriter, r *http.Request) {
		log.Debug(r.URL.Query())
		newGame := game.NewGame(cfg)
		// Save to DB
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := db.Insert(ctx, newGame.GameID, newGame.Board)
		if err != nil {
			j, err := newGame.Board.Json()
			if err != nil {
				log.Error("Error convering GameBoard to JSON", log.Fields{
					"gameID": newGame.GameID,
					"Board":  newGame.Board.String(),
					"error":  err,
				})
			}
			log.Error("Error writing GameBoard to DB", log.Fields{
				"gameID": newGame.GameID,
				"Board":  j,
				"error":  err,
			})
		}
		log.Debug("Created new Game", log.Fields{
			"gameID":    newGame.GameID,
			"GameBoard": newGame.Board.String(),
		})
		// Send joinGame event to EventRouter
		joinGameEvent := fmt.Sprintf(`{
			"type": "joinGame",
			"gameID": "%s",
			"sessionID": "%s",
			"timestamp": %d,
			"payload": {
				"status": "",
				"message": {
					"clientType": "server"
				}
			}
		}`, newGame.GameID, sessionID, time.Now().Unix())
		log.Debug("Writing joinGame event to EventRouter", log.Fields{
			"event": joinGameEvent,
		})
		err = ws.WriteMessage(websocket.TextMessage, []byte(joinGameEvent))
		if err != nil {
			log.Error("Error writing joinGame event to EventRouter", log.Fields{
				"gameID": newGame.GameID,
				"error":  err,
			})
		}

		// Send response
		_, err = w.Write([]byte(newGame.GameID))
		if err != nil {
			log.Error("Error responding to request to /newgame", log.Fields{
				"error": err,
			})
		}
	})

	log.Infof("Starting server at port %d", cfg.ServerPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ServerPort), nil); err != nil {
		log.Fatal(err)
	}
}

func connectToEventRouter(cfg config.Config) (*websocket.Conn, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   cfg.EventRouterURL,
		Path:   "/ws",
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

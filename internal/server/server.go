package server

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/CayenneLow/Codenames/internal/config"
	"github.com/CayenneLow/Codenames/internal/game"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func Start(cfg config.Config) {
	games := make(map[string](game.GameState))

	// Init connection to EventRouter
	ws, err := connectToEventRouter(cfg)
	if err != nil {
		log.Error("Error connecting to event router", log.Fields{
			"error": err,
		})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context, games map[string](game.GameState)) {
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
	}(ctx, games)

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

	http.HandleFunc("/newgame", func(w http.ResponseWriter, r *http.Request) {
		log.Debug(r.URL.Query())
		newGame := game.NewGame(cfg)
		games[newGame.GameID] = newGame
		_, err := w.Write([]byte(newGame.GameID))
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

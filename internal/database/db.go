package database

import (
	"context"

	"github.com/CayenneLow/Codenames/internal/config"
	"github.com/CayenneLow/Codenames/internal/game"
	redis "github.com/go-redis/redis/v9"
)

type Database interface {
	Disconnect(ctx context.Context) error
	ReadBoardByGameID(ctx context.Context, gameID string) (game.Board, error)
	Insert(ctx context.Context, gameID string, board game.Board) error
}

type database struct {
	dbClient     *redis.Client
	dbName       string
	dbCollection string
}

func Init(ctx context.Context, cfg config.Config) Database {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.DbURI,
		Password: "",
		DB:       0,
	})

	db := database{
		dbClient: client,
	}

	return &db
}

func (d *database) Disconnect(ctx context.Context) error {
	return d.dbClient.Close()
}

func (d *database) ReadBoardByGameID(ctx context.Context, gameID string) (game.Board, error) {
	res, err := d.dbClient.Get(ctx, gameID).Result()
	if err != nil {
		return game.Board{}, nil
	}
	return game.FromJson([]byte(res))
}

func (d *database) Insert(ctx context.Context, gameID string, board game.Board) error {
	b, err := board.Json()
	if err != nil {
		return err
	}
	_, err = d.dbClient.Set(ctx, gameID, []byte(b), 0).Result()
	return err
}

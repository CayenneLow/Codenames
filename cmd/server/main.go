package main

import (
	"context"
	"time"

	"github.com/CayenneLow/Codenames/internal/config"
	"github.com/CayenneLow/Codenames/internal/database"
	"github.com/CayenneLow/Codenames/internal/server"
)

func main() {
	cfg := config.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db := database.Init(ctx, cfg)
	server.Start(cfg, db)
}

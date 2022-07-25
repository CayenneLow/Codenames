package main

import (
	"github.com/CayenneLow/Codenames/internal/config"
	"github.com/CayenneLow/Codenames/internal/server"
)

func main() {
	cfg := config.Init()
	server.Start(cfg)
}

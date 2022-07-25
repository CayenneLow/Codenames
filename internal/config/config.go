package config

import (
	"log"
	"os"
	"strings"

	"github.com/CayenneLow/Codenames/internal/logger"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	EventRouterURL string `envconfig:"EVENT_ROUTER_URL" default:"localhost:8080"`
	ServerPort     int    `envconfig:"SERVER_PORT" default:"8080"`
	LogLevel       string `envconfig:"LOG_LEVEL" default:"INFO"`

	// Game Configs
	Words              []string
	Wordfilepath       string `envconfig:"WORD_FILEPATH"`
	BoardSize          int    `envconfig:"BOARD_SIZE"`
	BoardNRow          int    `envconfig:"BOARD_N_ROW"`
	BoardNCol          int    `envconfig:"BOARD_N_COL"`
	NTeams             int    `envconfig:"N_TEAMS"`
	NGuessStartingTeam int    `envconfig:"N_GUESS_STARTING_TEAM"`
	NGuessOtherTeams   int    `envconfig:"N_GUESS_OTHER_TEAM"`
	NGuessTotal        int    `envconfig:"N_GUESS_TOTAL"`
	NDeath             int    `envconfig:"N_DEATH"`
}

func Init() Config {
	cfg := Config{}
	// Initialize ENV
	err := envconfig.Process("", &cfg)
	initWords(&cfg)
	// Initialize logger
	logger.Init(cfg.LogLevel)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error initializing ENV variables"))
	}
	return cfg
}

func initWords(cfg *Config) {
	data, _ := os.ReadFile(cfg.Wordfilepath)
	cfg.Words = strings.Split(string(data), "\n")
}

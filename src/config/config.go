package config

import (
	"os"
	"strings"
)

type Config struct {
	Words              []string
	BoardSize          int
	BoardNRow          int
	BoardNCol          int
	NTeams             int
	NGuessStartingTeam int
	NGuessOtherTeams   int
	NGuessTotal        int
	NDeath             int
}

var Configuration Config

func Init() {
	Configuration = Config{}
	initWords()
}

func initWords() {
	data, _ := os.ReadFile("./src/assets/words.txt")
	Configuration.Words = strings.Split(string(data), "\n")
	Configuration.BoardNRow = 5
	Configuration.BoardNCol = 5
	Configuration.NTeams = 2
	Configuration.BoardSize = Configuration.BoardNRow * Configuration.BoardNCol
	Configuration.NGuessStartingTeam = 9
	Configuration.NGuessOtherTeams = 8
	Configuration.NGuessTotal = Configuration.NGuessStartingTeam + Configuration.NGuessOtherTeams
	Configuration.NDeath = 1
}

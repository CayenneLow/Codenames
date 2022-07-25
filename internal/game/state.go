package game

import (
	"math/rand"
	"strings"

	"github.com/CayenneLow/Codenames/internal/config"
	"github.com/CayenneLow/Codenames/internal/game/enum"
	"github.com/google/uuid"
	"github.com/meirf/gopart"
	log "github.com/sirupsen/logrus"
)

type GameState struct {
	Config    config.Config
	GameID    string
	ClientIDs []uint32
	CurrTeam  enum.Team
	NGuess    int               // The number of guesses the current team has left
	Remains   map[enum.Team]int // The number of cards remaining for each team
	Board     Board
}

func NewGame(cfg config.Config) GameState {
	board, startingTeam := generateBoard(cfg)
	remains := make(map[enum.Team](int))
	remains[startingTeam] = cfg.NGuessStartingTeam
	remains[startingTeam.Opposite()] = cfg.NGuessOtherTeams

	gameState := GameState{
		Config:    cfg,
		GameID:    newGameId(),
		ClientIDs: make([]uint32, 0),
		CurrTeam:  startingTeam,
		NGuess:    -1,
		Remains:   remains,
		Board:     board,
	}
	return gameState
}

func newGameId() string {
	newUuid := uuid.NewString()
	gameID := strings.ToUpper(newUuid[:5])
	return gameID
}

func (gs *GameState) Guess(row int, col int, team enum.Team) {
	log.Debugf("Guessing cell: %v, %v for Game: %v", row, col, gs.GameID)
	cell := &gs.Board.Cells[row][col]
	cell.Guessed = true
	if cell.Team == team {
		// TODO: A lot of work to be done, need to define a bunch of handlers
		// Correct guess
		gs.NGuess -= 1
		gs.Remains[team] -= 1

	}
}

// Returns board object and starting team
func generateBoard(cfg config.Config) (Board, enum.Team) {
	// Get 25 words
	wordIndexes := make(map[string]bool) // a map so this operation can remain O(n), value is not used
	for len(wordIndexes) < cfg.BoardSize {
		num := rand.Intn(len(cfg.Words))
		word := cfg.Words[num]
		if _, ok := wordIndexes[word]; !ok {
			// word doesn't exist, insert new entry
			wordIndexes[word] = true // we don't care about value
		}
	}
	// Map to Cells
	cells := make([]Cell, cfg.BoardSize)
	i := 0
	for key := range wordIndexes {
		cell := Cell{
			Word:    key,
			Team:    enum.NEUTRAL_TEAM,
			Guessed: false,
		}
		cells[i] = cell
		i++
	}
	// Partition into BoardNRow x BoardNCol
	cellGrid := make([][]Cell, cfg.BoardNRow)
	i = 0
	for idx := range gopart.Partition(len(cells), cfg.BoardNCol) {
		cellGrid[i] = cells[idx.Low:idx.High]
		i++
	}

	// Decide starting startingTeam
	startingTeam := enum.Team(rand.Intn(cfg.NTeams))
	// Decide on 9 words for starting team
	assignTeamToCells(cfg, cfg.NGuessStartingTeam, cellGrid, startingTeam)
	// Decide on 8 words for other team
	assignTeamToCells(cfg, cfg.NGuessOtherTeams, cellGrid, startingTeam.Opposite())
	// Decide on Death word
	assignTeamToCells(cfg, cfg.NDeath, cellGrid, enum.DEATH_TEAM)

	board := Board{Cells: cellGrid}
	return board, startingTeam
}

func assignTeamToCells(cfg config.Config, nWords int, cellGrid [][]Cell, team enum.Team) {
	i := 0
	for i < nWords {
		cellRow, cellCol := getRandCell(cfg)
		if cellGrid[cellRow][cellCol].Team == enum.NEUTRAL_TEAM {
			cellGrid[cellRow][cellCol].Team = team
			i++
		}
	}
}

func getRandCell(cfg config.Config) (int, int) {
	cellRow := rand.Intn(cfg.BoardNRow)
	cellCol := rand.Intn(cfg.BoardNCol)
	return cellRow, cellCol
}

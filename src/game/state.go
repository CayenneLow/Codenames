package game

import (
	"math/rand"

	"github.com/CayenneLow/Codenames/src/config"
	"github.com/CayenneLow/Codenames/src/game/enum"
	"github.com/CayenneLow/Codenames/src/logger"
	"github.com/google/uuid"
	"github.com/meirf/gopart"
)

type GameState struct {
	GameID    uuid.UUID
	ClientIDs []uint32
	CurrTeam  enum.Team
	NGuess    int               // The number of guesses the current team has left
	Remains   map[enum.Team]int // The number of cards remaining for each team
	Board     Board
}

func NewGame() GameState {
	newUuid := uuid.New()
	board, startingTeam := generateBoard()
	remains := make(map[enum.Team](int))
	remains[startingTeam] = config.Configuration.NGuessStartingTeam
	remains[startingTeam.Opposite()] = config.Configuration.NGuessOtherTeams

	gameState := GameState{
		GameID:    newUuid,
		ClientIDs: make([]uint32, 0),
		CurrTeam:  startingTeam,
		NGuess:    -1,
		Remains:   remains,
		Board:     board,
	}
	return gameState
}

func (gs *GameState) Guess(row int, col int, team enum.Team) {
	logger.Debugf("Guessing cell: %v, %v for Game: %v", row, col, gs.GameID)
	cell := &gs.Board.Cells[row][col]
	cell.guessed = true
	if cell.team == team {
		// TODO: A lot of work to be done, need to define a bunch of handlers
		// Correct guess
		gs.NGuess -= 1
		gs.Remains[team] -= 1

	}
}

// Returns board object and starting team
func generateBoard() (Board, enum.Team) {
	// Get 25 words
	wordIndexes := make(map[string]bool) // a map so this operation can remain O(n), value is not used
	for len(wordIndexes) < config.Configuration.BoardSize {
		num := rand.Intn(len(config.Configuration.Words))
		word := config.Configuration.Words[num]
		if _, ok := wordIndexes[word]; !ok {
			// word doesn't exist, insert new entry
			wordIndexes[word] = true // we don't care about value
		}
	}
	// Map to Cells
	cells := make([]Cell, config.Configuration.BoardSize)
	i := 0
	for key := range wordIndexes {
		cell := Cell{
			word:    key,
			team:    enum.NEUTRAL_TEAM,
			guessed: false,
		}
		cells[i] = cell
		i++
	}
	// Partition into BoardNRow x BoardNCol
	cellGrid := make([][]Cell, config.Configuration.BoardNRow)
	i = 0
	for idx := range gopart.Partition(len(cells), config.Configuration.BoardNCol) {
		cellGrid[i] = cells[idx.Low:idx.High]
		i++
	}

	// Decide starting startingTeam
	startingTeam := enum.Team(rand.Intn(config.Configuration.NTeams))
	// Decide on 9 words for starting team
	assignTeamToCells(config.Configuration.NGuessStartingTeam, cellGrid, startingTeam)
	// Decide on 8 words for other team
	assignTeamToCells(config.Configuration.NGuessOtherTeams, cellGrid, startingTeam.Opposite())
	// Decide on Death word
	assignTeamToCells(config.Configuration.NDeath, cellGrid, enum.DEATH_TEAM)

	board := Board{Cells: cellGrid}
	logger.Debug(board.String())
	return board, startingTeam
}

func assignTeamToCells(nWords int, cellGrid [][]Cell, team enum.Team) {
	i := 0
	for i < nWords {
		cellRow, cellCol := getRandCell()
		if cellGrid[cellRow][cellCol].team == enum.NEUTRAL_TEAM {
			cellGrid[cellRow][cellCol].team = team
			i++
		}
	}
}

func getRandCell() (int, int) {
	cellRow := rand.Intn(config.Configuration.BoardNRow)
	cellCol := rand.Intn(config.Configuration.BoardNCol)
	return cellRow, cellCol
}

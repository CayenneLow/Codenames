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
	NGuess    int
	CurrTeam  enum.Team
	Board     Board
}

func NewGame() GameState {
	newUuid := uuid.New()
	board, startingTeam := generateBoard()

	// for _, col := range board.Cells {
	// 	for _, cell := range col {
	// 		fmt.Print(cell.String() + " ")
	// 	}
	// 	fmt.Println("")
	// }

	gameState := GameState{
		GameID:    newUuid,
		ClientIDs: make([]uint32, 0),
		NGuess:    -1,
		CurrTeam:  startingTeam,
		Board:     board,
	}
	return gameState
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
	assignTeamToCells(9, cellGrid, startingTeam)
	// Decide on 8 words for other team
	assignTeamToCells(8, cellGrid, startingTeam.Opposite())
	// Decide on Death word
	assignTeamToCells(1, cellGrid, enum.DEATH_TEAM)

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

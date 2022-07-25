package game

import (
	"fmt"

	"github.com/CayenneLow/Codenames/internal/game/enum"
)

type Cell struct {
	Word    string    `json:"word"`
	Team    enum.Team `json:"team"`
	Guessed bool      `json:"guessed"`
}

func (c Cell) String() string {
	return fmt.Sprintf("%10s (%7s)", c.Word, c.Team.String())
}

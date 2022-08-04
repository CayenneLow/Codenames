package game

import (
	"fmt"
)

type Cell struct {
	Word    string `json:"word"`
	Team    string `json:"team"`
	Guessed bool   `json:"guessed"`
}

func (c Cell) String() string {
	return fmt.Sprintf("%10s (%7s)", c.Word, c.Team)
}

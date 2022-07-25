package game

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Board struct {
	Cells [][]Cell `json:"cells"`
}

func (b Board) String() string {
	var sb strings.Builder

	sb.WriteString("\n")
	cellGrid := b.Cells
	for _, row := range cellGrid {
		for _, col := range row {
			sb.WriteString(fmt.Sprintf("%s ||", col.String()))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b Board) Json() ([]byte, error) {
	return json.Marshal(b)
}

func FromJson(j []byte) (Board, error) {
	var board Board
	err := json.Unmarshal(j, &board)
	return board, err
}

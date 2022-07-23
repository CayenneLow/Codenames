package game

import (
	"fmt"
	"strings"
)

type Board struct {
	Cells [][]Cell
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

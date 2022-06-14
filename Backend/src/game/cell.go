package game

import (
	"fmt"

	"github.com/CayenneLow/Codenames/src/game/enum"
)

type Cell struct {
	word    string
	team    enum.Team
	guessed bool
}

func (c Cell) String() string {
	return fmt.Sprintf("%10s (%7s)", c.word, c.team.String())
}

package enum

type Team int

const (
	RED_TEAM Team = iota
	BLUE_TEAM
	NEUTRAL_TEAM
	DEATH_TEAM
)

func (t Team) String() string {
	switch t {
	case RED_TEAM:
		return "Red"
	case BLUE_TEAM:
		return "Blue"
	case NEUTRAL_TEAM:
		return "Neutral"
	case DEATH_TEAM:
		return "Death"
	default:
		return "Unknown"
	}
}

func (t Team) Opposite() Team {
	switch t {
	case RED_TEAM:
		return BLUE_TEAM
	case BLUE_TEAM:
		return RED_TEAM
	default:
		return -1 // Unknown
	}
}

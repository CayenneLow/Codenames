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
		return "RED"
	case BLUE_TEAM:
		return "BLUE"
	case NEUTRAL_TEAM:
		return "NEUTRAL"
	case DEATH_TEAM:
		return "DEATH"
	default:
		return "UNKNOWN"
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

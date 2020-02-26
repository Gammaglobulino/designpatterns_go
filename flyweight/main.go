package flyweight

import "time"

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte //photo
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type TeamFlyWeightFactory struct {
	CreatedTeams map[int]*Team
}

func (t *TeamFlyWeightFactory) GetTeam(teamId int) *Team {
	if t.CreatedTeams[teamId] != nil {
		return t.CreatedTeams[teamId]
	}
	team := getTeamFactory(teamId)
	t.CreatedTeams[teamId] = &team
	return t.CreatedTeams[teamId]
}
func (t *TeamFlyWeightFactory) GetNumberOfObjects() int {
	return len(t.CreatedTeams)
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_A:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TEAM:B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

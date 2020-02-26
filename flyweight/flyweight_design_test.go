package flyweight

import (
	"../flyweight"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTeamFlyWeightFactory_GetTeam(t *testing.T) {
	factory := flyweight.TeamFlyWeightFactory{make(map[int]*flyweight.Team, 0)}
	teamA1 := factory.GetTeam(flyweight.TEAM_A)
	assert.NotNil(t, teamA1, "The pointer to TEAM_A was nil")
	teamA2 := factory.GetTeam(flyweight.TEAM_A)
	assert.NotNil(t, teamA2, "The pointer to TEAM_A was nil")
	assert.Equal(t, teamA1, teamA2, "TEAM_A pointers were not the same")
	assert.Equal(t, factory.GetNumberOfObjects(), 1, fmt.Sprintf("Number of object created not 1, actual: %d", factory.GetNumberOfObjects()))
}

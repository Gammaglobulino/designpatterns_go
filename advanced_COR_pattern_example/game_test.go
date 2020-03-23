package advanced_COR_pattern_example

import (
	"../advanced_COR_pattern_example"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestNewCreature(t *testing.T) {
	game := &advanced_COR_pattern_example.Game{sync.Map{}}
	goblin := advanced_COR_pattern_example.NewCreature(game, "Goblin", 1, 1)
	assert.EqualValues(t, "Goblin (1/1)", goblin.String())
}

func TestNewDoubleAttackModifier(t *testing.T) {
	game := &advanced_COR_pattern_example.Game{sync.Map{}}
	goblin := advanced_COR_pattern_example.NewCreature(game, "Goblin", 1, 1)
	advanced_COR_pattern_example.NewDoubleAttackModifier(game, goblin)
	assert.EqualValues(t, "Goblin (2/1)", goblin.String())
}

func TestNewDoubleAttackModifierSub_Unsub(t *testing.T) {
	game := &advanced_COR_pattern_example.Game{sync.Map{}}
	goblin := advanced_COR_pattern_example.NewCreature(game, "Goblin", 1, 1)
	d := advanced_COR_pattern_example.NewDoubleAttackModifier(game, goblin)
	assert.EqualValues(t, "Goblin (2/1)", goblin.String())
	d.Close()
	assert.EqualValues(t, "Goblin (1/1)", goblin.String())
}

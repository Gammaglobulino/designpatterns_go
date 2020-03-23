package chain_of_responsability_method_chain

import (
	"../chain_of_responsability_method_chain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCreature(t *testing.T) {
	goblin := chain_of_responsability_method_chain.NewCreature("Goblin", 1, 1)
	assert.EqualValues(t, "Goblin (1/1)", goblin.String())
}

func TestNewCreatureModifier(t *testing.T) {
	goblin := chain_of_responsability_method_chain.NewCreature("Goblin", 1, 1)
	root := chain_of_responsability_method_chain.NewCreatureModifier(goblin)
	root.Add(chain_of_responsability_method_chain.NewDobuleAttackModifier(goblin))
	root.Handle()
	assert.EqualValues(t, "Goblin (2/1)", goblin.String())
}

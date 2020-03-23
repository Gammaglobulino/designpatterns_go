package chain_of_responsability_method_chain

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{
		Name:    name,
		Attack:  attack,
		Defense: defense,
	}
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}
type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

func (cm *CreatureModifier) Add(m Modifier) {
	if cm.next != nil {
		cm.Add(m)
	} else {
		cm.next = m
	}
}
func (cm *CreatureModifier) Handle() {
	if cm.next != nil {
		cm.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDobuleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{
		creature: c,
	}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling ", d.creature.Name, "\b's attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()

}

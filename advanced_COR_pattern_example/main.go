package advanced_COR_pattern_example

import (
	"fmt"
	"sync"
)

//CoR,Mediator,Observer,CQS (Command Query Separator)

type Argument int

const (
	Attack Argument = iota
	Defense
)

// used to send a message to Observers..
type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(q *Query)
}
type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	Observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.Observers.Store(o, struct{}{})
}
func (g *Game) Unsubscribe(o Observer) {
	g.Observers.Delete(o)
}
func (g *Game) Fire(q *Query) {
	g.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type Creature struct {
	game            *Game //mediator
	Name            string
	attack, defense int
}

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
	return &Creature{
		game:    game,
		Name:    name,
		attack:  attack,
		defense: defense,
	}
}

func (c *Creature) Attack() int {
	q := Query{
		CreatureName: c.Name,
		WhatToQuery:  Attack,
		Value:        c.attack,
	}
	c.game.Fire(&q)
	return q.Value
}
func (c *Creature) Defense() int {
	q := Query{
		CreatureName: c.Name,
		WhatToQuery:  Defense,
		Value:        c.defense,
	}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(), c.Defense())
}

type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func (c *CreatureModifier) Handle(q *Query) {

}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{CreatureModifier{
		game:     g,
		creature: c,
	}}
	g.Subscribe(d)
	return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == d.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

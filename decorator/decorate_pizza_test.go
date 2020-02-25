package decorator

import (
	"../decorator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &decorator.PizzaDecorator{}
	pizzaNext, err := pizza.AddIngredient()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, "Pizza with the following ingredients:", pizzaNext)
}
func TestOnion_AddIngredient(t *testing.T) {
	onion := &decorator.Onion{}
	onionNext, err := onion.AddIngredient()
	assert.NotNil(t, err, "Ingredient added with no base")
	assert.EqualValues(t, "", onionNext)
	onion = &decorator.Onion{&decorator.PizzaDecorator{}}
	onionNext, err = onion.AddIngredient()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, "Pizza with the following ingredients: onion", onionNext)
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &decorator.Meat{}
	meatNext, err := meat.AddIngredient()
	assert.NotNil(t, err, "Ingredient added with no base")
	assert.EqualValues(t, "", meatNext)
	meat = &decorator.Meat{&decorator.PizzaDecorator{}}
	meatNext, err = meat.AddIngredient()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, "Pizza with the following ingredients: meat", meatNext)
}

func TestPizzaDecorator_Full(t *testing.T) {
	pizza := &decorator.Onion{&decorator.Meat{&decorator.PizzaDecorator{}}}
	pizzaNext, err := pizza.AddIngredient()
	if err != nil {
		t.Fatal()
	}
	assert.EqualValues(t, "Pizza with the following ingredients: meat onion", pizzaNext)
}

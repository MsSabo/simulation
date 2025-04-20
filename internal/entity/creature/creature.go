package creature

import (
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/entity"
	"github.com/MsSabo/simulation/internal/gameboard"
)

type Animal interface {
	entity.Entity
	Eat(gameboard *gameboard.Gameboard, cell internal.Cell)
	Move(gameboard *gameboard.Gameboard, cell internal.Cell) bool
}

type AnimalsCoordinates struct {
	Cell   internal.Cell
	Animal Animal
}

func GetAllAnimals(gb *gameboard.Gameboard) []AnimalsCoordinates {
	var animals []AnimalsCoordinates

	for cell, val := range gb.Board {
		switch val.(type) {
		case *Sheep:
			animals = append(animals, AnimalsCoordinates{cell, val.(Animal)})
		case *Predator:
			animals = append(animals, AnimalsCoordinates{cell, val.(Animal)})
		default:
		}
	}

	return animals
}

type animalParam struct {
	// @todo добавить различные параметры (скорость, штраф за голод и тп)
	speed    int
	health   int
	calories int
}

func (state *animalParam) DecreaseHealth(damage int) {
	state.health -= damage
}

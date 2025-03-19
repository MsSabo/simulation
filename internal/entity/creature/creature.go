package creature

import (
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/entity"
	"github.com/MsSabo/simulation/internal/gameboard"
)

type Animal interface {
	entity.Entity
	Eat(gameboard *gameboard.Gameboard, cell internal.Cell)
	Move(gameboard *gameboard.Gameboard) bool
}

func GetAllAnimals(gb *gameboard.Gameboard) []Animal {
	var animals []Animal

	for _, val := range gb.Board {
		switch val.(type) {
		case *Sheep:
			animals = append(animals, val.(Animal))
		case *Predator:
			animals = append(animals, val.(Animal))
		default:
		}
	}

	return animals
}

type animalParam struct {
	// @todo добавить различные параметры (скорость, штраф за голод и тп)
}

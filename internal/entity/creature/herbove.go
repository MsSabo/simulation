package creature

import (
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/entity"
	"github.com/MsSabo/simulation/internal/gameboard"
	pathfinder "github.com/MsSabo/simulation/internal/routebuilder"
)

var sheepRoutebuilder = pathfinder.NewBfsRouteBuilder[*entity.Grass]()

type Sheep struct {
	animalParam
}

func (s *Sheep) GetSign() string {
	return entity.SheepSigns
}

func (s *Sheep) Eat(gameboard *gameboard.Gameboard, cell internal.Cell) {
	gameboard.RemoveEntity(cell)
}

func (s *Sheep) Move(gameboard *gameboard.Gameboard, start internal.Cell) bool {
	route := sheepRoutebuilder.FindRoute(*gameboard, start)

	if len(route) == 0 {
		return false
	}

	if len(route) == 1 {
		s.Eat(gameboard, route[0])
	} else {
		gameboard.Move(start, route[0])
	}

	return true
}

func NewSheep(x, y int) *Sheep {
	s := Sheep{animalParam{speed: 1, health: 15, calories: 0}}
	return &s
}

package creature

import (
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/entity"
	"github.com/MsSabo/simulation/internal/gameboard"
	pathfinder "github.com/MsSabo/simulation/internal/routebuilder"
)

var predatorRoutebuilder = pathfinder.NewBfsRouteBuilder[*Sheep]()

type Predator struct {
	animalParam
	attack int
}

func (s *Predator) GetSign() string {
	return entity.FoxSigns
}

func (s *Predator) Eat(gameboard *gameboard.Gameboard, cell internal.Cell) {
	sheep, ok := gameboard.GetEntity(cell).(*Sheep)
	if !ok {
		panic("Fox can't get entity")
	}

	sheep.DecreaseHealth(s.attack)
	if sheep.health <= 0 {
		gameboard.RemoveEntity(cell)
	}
}

func (s *Predator) Move(gameboard *gameboard.Gameboard, start internal.Cell) bool {
	route := predatorRoutebuilder.FindRoute(*gameboard, start)
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

func NewPredator(x, y int) *Predator {
	s := Predator{animalParam{}, 5}
	return &s
}

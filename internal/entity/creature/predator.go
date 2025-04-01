package creature

import (
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/gameboard"
	pathfinder "github.com/MsSabo/simulation/internal/routebuilder"
)

var predatorRoutebuilder = pathfinder.NewBfsRouteBuilder[*Sheep]()

type Predator struct {
	internal.Cell
	animalParam
}

func (s *Predator) GetSign() string {
	return "🦊"
}

func (s *Predator) Eat(gameboard *gameboard.Gameboard, cell internal.Cell) {
	gameboard.RemoveEntity(cell)
}

func (s *Predator) Move(gameboard *gameboard.Gameboard) bool {
	route := predatorRoutebuilder.FindRoute(*gameboard, s.Cell)
	if len(route) == 0 {
		return false
	}

	if len(route) == 1 {
		s.Eat(gameboard, route[0])
	} else {
		gameboard.Move(s.Cell, route[0])
		s.Set(route[0].Get())
	}

	return true
}

func NewPredator(x, y int) *Predator {
	s := Predator{internal.MakeCell(x, y), animalParam{}}
	return &s
}

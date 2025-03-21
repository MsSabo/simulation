package creature

import (
	"fmt"
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/entity"
	"github.com/MsSabo/simulation/internal/gameboard"
	pathfinder "github.com/MsSabo/simulation/internal/routebuilder"
)

var sheepRoutebuilder = pathfinder.NewBfsRouteBuilder[*entity.Grass]()

type Sheep struct {
	internal.Cell
	animalParam
}

func (s *Sheep) GetSign() string {
	return "🐑"
}

func (s *Sheep) Eat(gameboard *gameboard.Gameboard, cell internal.Cell) {
	gameboard.RemoveEntity(cell)
	fmt.Println("Sheep eated grass")
}

func (s *Sheep) Move(gameboard *gameboard.Gameboard) bool {
	route := sheepRoutebuilder.FindRoute(*gameboard, s.Cell)
	if len(route) == 0 {
		return false
	}

	if len(route) == 1 {
		s.Eat(gameboard, route[0])
	} else {
		gameboard.Move(s.Cell, route[0])
		s.Set(route[0].Get())
		fmt.Println("Sheep move")
	}

	return true
}

func NewSheep(x, y int) *Sheep {
	s := Sheep{internal.MakeCell(x, y), animalParam{}}
	return &s
}

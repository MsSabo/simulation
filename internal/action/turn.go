package action

import (
	"github.com/MsSabo/simulation/internal/entity/creature"
	"github.com/MsSabo/simulation/internal/gameboard"
)

type MoveAll struct{}

func (m *MoveAll) Make(g *gameboard.Gameboard) {
	animals := creature.GetAllAnimals(g)
	for _, obj := range animals {
		obj.Animal.Move(g, obj.Cell)
	}
}

func MakeMoveAll() Action {
	return &MoveAll{}
}

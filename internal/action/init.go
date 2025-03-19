package action

import (
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/gameboard"
)

type initAction struct {
}

func (a *initAction) Make(g gameboard.Gameboard) {
	for i := range g.GetX() {
		for j := range g.GetY() {
			ent := randomEntity()
			if ent != nil {
				ent.Set(i, j)
				g.AddEntity(internal.MakeCell(i, j), ent)
			}
		}
	}
}

func MakeInitAction() Action {
	return &initAction{}
}

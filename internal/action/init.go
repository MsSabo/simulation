package action

import (
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/gameboard"
	"math/rand"
	"time"
)

type initAction struct {
}

func (a *initAction) Make(g gameboard.Gameboard) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range g.GetX() {
		for j := range g.GetY() {
			ent := randomEntity(rand.Float32())
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

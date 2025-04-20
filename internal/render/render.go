package render

import (
	"fmt"

	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/gameboard"
	"github.com/mattn/go-runewidth"
)

type Render interface {
	PrintGameBoard(gameboard gameboard.Gameboard)
}

type renderImpl struct {
}

func alignSign(sign string) string {
	width := runewidth.StringWidth(sign)

	if width%2 != 0 {
		return sign + "."
	}

	return sign
}

func (r *renderImpl) PrintGameBoard(g gameboard.Gameboard) {
	xlength := g.Row()
	ylength := g.Columnt()

	for i := range xlength {
		for j := range ylength {
			if elem := g.GetEntity(internal.MakeCell(i, j)); elem != nil {
				fmt.Print(alignSign(elem.GetSign()))
			} else {
				fmt.Print("..")
			}
		}
		fmt.Println()
	}
}

func NewRender() Render {
	return &renderImpl{}
}

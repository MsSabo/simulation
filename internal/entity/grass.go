package entity

import "github.com/MsSabo/simulation/internal"

type Grass struct {
	internal.Cell
}

func (g *Grass) GetSign() string {
	return GrassSigns
}

func NewGrass(x, y int) *Grass {
	g := Grass{internal.MakeCell(x, y)}
	return &g
}

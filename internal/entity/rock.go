package entity

import "github.com/MsSabo/simulation/internal"

type Rock struct {
	internal.Cell
}

func (g *Rock) GetSign() string {
	return "⛰️"
}

func NewRock(x, y int) *Rock {
	g := Rock{internal.MakeCell(x, y)}
	return &g
}

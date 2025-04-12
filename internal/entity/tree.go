package entity

import "github.com/MsSabo/simulation/internal"

type Tree struct {
	internal.Cell
}

func (g *Tree) GetSign() string {
	return TreeSigns
}

func NewTree(x, y int) *Tree {
	g := Tree{internal.MakeCell(x, y)}
	return &g
}

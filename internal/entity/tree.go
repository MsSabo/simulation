package entity

type Tree struct {
}

func (g *Tree) GetSign() string {
	return TreeSigns
}

func NewTree(x, y int) *Tree {
	g := Tree{}
	return &g
}

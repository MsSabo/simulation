package entity

type Grass struct {
}

func (g *Grass) GetSign() string {
	return GrassSigns
}

func NewGrass(x, y int) *Grass {
	g := Grass{}
	return &g
}

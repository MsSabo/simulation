package entity

type Rock struct {
}

func (g *Rock) GetSign() string {
	return RockSigns
}

func NewRock(x, y int) *Rock {
	g := Rock{}
	return &g
}

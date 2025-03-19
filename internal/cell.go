package internal

type Cell struct {
	x int
	y int
}

type ICell interface{
	Set(x, y int)
	Get() (x, y int)
	Plus(Cell) Cell
}

func(c *Cell) Set(x, y int) {
	c.x = x
	c.y = y
}

func(c *Cell) Get() (x, y int) {
	return c.x, c.y
}

func(c *Cell) Plus(p Cell) Cell {
	res := MakeCell(c.x + p.x, c.y + p.y)
	return res
}

func NewCell(x, y int) *Cell {
	c := Cell{x, y}
	return &c
}

func MakeCell(x, y int) Cell {
	return Cell{x, y}
}
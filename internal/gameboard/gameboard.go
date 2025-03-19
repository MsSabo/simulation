package gameboard

import (
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/entity"
)

type Gameboard struct {
	xlength int
	ylength int
	Board   map[internal.Cell]entity.Entity // @todo скрыть
}

func NewGameboard(x, y int) Gameboard {
	return Gameboard{x, y, make(map[internal.Cell]entity.Entity)}
}

func (gameboard *Gameboard) AddEntity(c internal.Cell, e entity.Entity) {
	gameboard.Board[c] = e
}

func (gameboard *Gameboard) RemoveEntity(c internal.Cell) {
	delete(gameboard.Board, c)
}

func (gameboard *Gameboard) Move(old internal.Cell, new internal.Cell) {
	entity := gameboard.Board[old]
	delete(gameboard.Board, old)
	gameboard.Board[new] = entity
}

func (gameboard *Gameboard) GetEntity(c internal.Cell) entity.Entity {
	if elem, ok := gameboard.Board[c]; ok {
		return elem
	}
	return nil
}

func (gameboard Gameboard) CellIsValid(c internal.Cell) bool {
	if x, y := c.Get(); x >= 0 && x < gameboard.xlength && y >= 0 && y < gameboard.ylength {
		return true
	}

	return false
}

func (gameboard Gameboard) GetX() int {
	return gameboard.xlength
}

func (gameboard Gameboard) GetY() int {
	return gameboard.ylength
}

package gameboard

import (
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/entity"
)

type Gameboard struct {
	row    int
	column int
	Board  map[internal.Cell]entity.Entity // @todo скрыть
}

func NewGameboard(row, column int) Gameboard {
	return Gameboard{row, column, make(map[internal.Cell]entity.Entity)}
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
	if x, y := c.Get(); x >= 0 && x < gameboard.row && y >= 0 && y < gameboard.column {
		return true
	}

	return false
}

func (gameboard Gameboard) IsFree(c internal.Cell) bool {
	return gameboard.Board[c] == nil
}

func (gameboard Gameboard) Row() int {
	return gameboard.row
}

func (gameboard Gameboard) Columnt() int {
	return gameboard.column
}

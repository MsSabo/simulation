package pathfinder

import (
	"github.com/MsSabo/simulation/internal/gameboard"
	"reflect"

	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/entity"
)

type GameObject interface {
	entity.Entity
}

type RouteBuilder[Target GameObject] interface {
	FindRoute(gameboard gameboard.Gameboard, start internal.Cell) []internal.Cell
}

type bfsRouteBuilder[Target GameObject] struct {
	t Target
}

func NewBfsRouteBuilder[Target GameObject]() RouteBuilder[Target] {
	return &bfsRouteBuilder[Target]{}
}

func (builder *bfsRouteBuilder[Target]) FindRoute(gameboard gameboard.Gameboard, start internal.Cell) []internal.Cell {
	directions := [4]internal.Cell{internal.MakeCell(1, 0), internal.MakeCell(0, 1),
		internal.MakeCell(-1, 0), internal.MakeCell(0, -1)}

	route := make(map[internal.Cell]internal.Cell, 0)

	queue := make([]internal.Cell, 0)
	queue = append(queue, start)
	var destination internal.Cell

	for len(queue) > 0 {
		destination = queue[0]
		queue = queue[1:]

		obj := (gameboard).GetEntity(destination)

		if reflect.TypeOf(builder.t) == reflect.TypeOf(obj) {
			break
		}

		for _, val := range directions {
			neighbor := destination.Plus(val)
			_, visited := route[neighbor]
			valid := (gameboard).CellIsValid(neighbor)
			neighobObj := (gameboard).GetEntity(neighbor)

			if (valid && !visited) && (neighobObj == nil || reflect.TypeOf(builder.t) == reflect.TypeOf(neighobObj)) {
				queue = append(queue, neighbor)
				route[neighbor] = destination
			}
		}
	}

	result := make([]internal.Cell, 0)
	if len(queue) == 0 {
		return result
	}

	result = append(result, destination)
	for node := route[destination]; node != start; node = route[node] {
		result = append(result, node)
	}

	// Reverse the new slice in place
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

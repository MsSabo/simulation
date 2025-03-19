package action

import (
	"github.com/MsSabo/simulation/internal/entity"
	"github.com/MsSabo/simulation/internal/entity/creature"
	"github.com/MsSabo/simulation/internal/gameboard"
	"reflect"
)

type Action interface {
	Make(gb gameboard.Gameboard)
}

// / Вынести это в интерфейс
func GameIsOver(gb gameboard.Gameboard) bool {
	sheep := creature.NewSheep(0, 0)

	sheepExists := isExists[*creature.Sheep](gb, sheep)

	return sheepExists
}

type gameObject interface {
	entity.Entity
}

func isExists[T gameObject](game gameboard.Gameboard, obj T) bool {
	for _, val := range game.Board {
		if reflect.TypeOf(obj) == reflect.TypeOf(val) {
			return true
		}
	}
	return false
}

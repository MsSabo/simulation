package action

import (
	"github.com/MsSabo/simulation/internal/entity"
	"github.com/MsSabo/simulation/internal/entity/creature"
)

func randomEntity(val float32) entity.Entity {
	if val <= 0.5 {
		return nil
	} else if val <= 0.7 {
		return entity.NewGrass(0, 0)
	} else if val <= 0.75 {
		return entity.NewRock(0, 0)
	} else if val <= 0.85 {
		return entity.NewTree(0, 0)
	} else if val <= 0.9 {
		return creature.NewSheep(0, 0)
	} else {
		return creature.NewPredator(0, 0)
	}
}

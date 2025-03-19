package action

import (
	"github.com/MsSabo/simulation/internal/entity"
	"github.com/MsSabo/simulation/internal/entity/creature"
	"math/rand"
	"time"
)

func randomEntity() entity.Entity {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomFloat := source.Float64()

	if randomFloat <= 0.85 {
		return nil
	} else if randomFloat <= 0.9 {
		return entity.NewGrass(0, 0)
	} else if randomFloat <= 0.94 {
		return entity.NewRock(0, 0)
	} else if randomFloat <= 0.98 {
		return entity.NewTree(0, 0)
	} else if randomFloat <= 0.99 {
		return creature.NewSheep(0, 0)
	} else {
		return creature.NewPredator(0, 0)
	}
}

package action

import (
	"fmt"
	"github.com/MsSabo/simulation/internal"
	"github.com/MsSabo/simulation/internal/config"
	"github.com/MsSabo/simulation/internal/entity"
	"github.com/MsSabo/simulation/internal/entity/creature"
	"github.com/MsSabo/simulation/internal/gameboard"
	"math/rand"
	"time"
)

type initAction struct {
}

func createEntity(name string, x, y int) entity.Entity {
	switch name {
	case config.Grass:
		return entity.NewGrass(x, y)
	case config.Tree:
		return entity.NewTree(x, y)
	case config.Rock:
		return entity.NewRock(x, y)
	case config.Sheep:
		return creature.NewSheep(x, y)
	case config.Fox:
		return creature.NewPredator(x, y)
	default:
		return nil
	}
}

func (a *initAction) Make(g *gameboard.Gameboard) {
	cfg, err := config.ParseConfig("../internal/config/config.xml")

	if err != nil {
		panic(err)
	}

	*g = gameboard.NewGameboard(cfg.MapSize.Xlength, cfg.MapSize.Ylength)
	fmt.Println("Size = ", g.GetX())
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, param := range cfg.Entity {
		fmt.Printf("Create %s count %d \n", param.Name, param.Quantity)
		for _ = range param.Quantity {
			var x, y int

			for ok := false; !ok; ok = g.IsFree(internal.MakeCell(x, y)) {
				x, y = rand.Intn(g.GetX()), rand.Intn(g.GetY())
			}
			ent := createEntity(param.Name, x, y)
			g.AddEntity(internal.MakeCell(x, y), ent)
		}
	}
}

func MakeInitAction() Action {
	return &initAction{}
}

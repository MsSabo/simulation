package simulation

import (
	"context"
	"fmt"
	"github.com/MsSabo/simulation/internal/action"
	"github.com/MsSabo/simulation/internal/gameboard"
	"github.com/MsSabo/simulation/internal/render"
	"sync"
	"time"
)

type Simulation struct {
	counter int
	render  render.Render
	gb      gameboard.Gameboard
	init    action.Action
	move    action.Action
}

func (s *Simulation) Init() {
	s.init.Make(s.gb)
	s.render.PrintGameBoard(s.gb)
	fmt.Println("Map created")
}

func (s *Simulation) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		isOver := action.GameIsOver(s.gb)
		for isOver {
			select {
			case <-ctx.Done():
				fmt.Println("context canceled")
				return
			default:
				s.move.Make(s.gb)
				isOver = action.GameIsOver(s.gb)
				s.counter++
				fmt.Printf("simulation step = %d, over = %v\n", s.counter, isOver)
				s.render.PrintGameBoard(s.gb)
				time.Sleep(time.Second * 3)
			}
		}
		fmt.Println("game over")
	}()
}

func NewSimulation(x, y int) *Simulation {
	gb := gameboard.NewGameboard(x, y)
	return &Simulation{render: render.NewRender(), gb: gb, init: action.MakeInitAction(), move: action.MakeMoveAll()}
}

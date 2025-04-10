package main

import (
	"context"
	"github.com/MsSabo/simulation/internal/simulation"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func main() {
	waitCh := make(chan struct{})

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	game := simulation.NewSimulation(5, 5)
	game.Init()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT|syscall.SIGTERM|syscall.SIGQUIT)

	game.Run(ctx, &wg)

	go func() {
		wg.Wait()
		close(waitCh)
	}()

	select {
	case <-c:
		cancel()
	case <-waitCh:
	}
}

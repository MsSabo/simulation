package main

import (
	"context"
	"github.com/MsSabo/simulation/internal/simulation"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wait sync.WaitGroup

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	game := simulation.NewSimulation(10, 15)
	game.Init()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT|syscall.SIGTERM|syscall.SIGQUIT)
	game.Run(ctx, &wait)

	<-c
	cancel()
	wait.Wait()
}

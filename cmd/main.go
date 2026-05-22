package main

import (
	"errors"
	"go_invaders/game"
)

func main() {
	var windowWidth int32 = 1920
	var windowHeight int32 = 1080

	g := game.New(windowWidth, windowHeight)
	g.Init()
	defer g.Destroy()

	if err := g.Run(); err != nil {
		panic(errors.New("failed to run the game: " + err.Error()))
	}
}

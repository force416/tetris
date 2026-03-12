package main

import (
	"log"

	"tetris/internal/app"
	"tetris/internal/game"
)

func main() {
	g := game.New(game.DefaultConfig(), nil)
	application := app.New(g)
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}

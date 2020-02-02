package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/ivorscott/packmango/packman"
)

func main() {
	g := packman.NewGame()

	if err := ebiten.Run(g.Update, g.ScreenWidth(), g.ScreenHeight(), 2, "Packman"); err != nil {
		log.Fatal(err)
	}
}

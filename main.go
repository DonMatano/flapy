package main

import (
	"github.com/DonMatano/flapy/config"
	"github.com/DonMatano/flapy/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(config.WindowWidth, config.WindowHeight, "Flapy bird")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	gameObject := game.GetGameInstance()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		gameObject.Loop()
		rl.EndDrawing()
	}
}

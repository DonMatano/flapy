package flapy

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	screenWidth  int32 = 800
	screenHeight int32 = 450
)

func loadTexture(path string) rl.Texture2D {
	texture2D := rl.LoadTexture(path)
	return texture2D
}

func Init() {
	rl.InitWindow(screenWidth, screenHeight, "flapy bird")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()
	bgTexture := loadTexture("assets/Background.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		rl.DrawTexture(bgTexture, screenHeight/2-bgTexture.Width/2, screenHeight/2-bgTexture.Height/2, rl.White)

		rl.EndDrawing()
	}
}

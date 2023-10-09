package flapy

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	characters "github.com/DonMatano/flapy/flapy/characters"
)

var (
	screenWidth  int32 = 1280
	screenHeight int32 = 800
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

	birdTexture := loadTexture("assets/bird/flying/frame-1.png")

	bird := characters.NewBird(characters.BirdParams{
		Length: 80, Width: 100, Position: rl.NewVector2(100, 100), Texture: birdTexture,
	})

	for !rl.WindowShouldClose() {
		bird.Update()
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		rl.DrawTexture(bgTexture, screenWidth/2-bgTexture.Width/2, screenHeight/2-bgTexture.Height/2, rl.White)
		// rl.DrawTexture(birdTexture, screenWidth/2-birdTexture.Width/2, screenHeight/2-birdTexture.Height/2, rl.White)
		bird.Draw()

		rl.EndDrawing()
	}
}

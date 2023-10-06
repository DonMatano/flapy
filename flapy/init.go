package flapy

import raylib "github.com/gen2brain/raylib-go/raylib"

var (
	screenWidth  int32 = 800
	screenHeight int32 = 450
)

func Init() {
	raylib.InitWindow(screenWidth, screenHeight, "flapy bird")
	raylib.SetTargetFPS(60)
	defer raylib.CloseWindow()
	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.White)
		raylib.DrawText("Hello To Flapy bird.", 100, 200, 16*2, raylib.Red)

		raylib.EndDrawing()
	}
}

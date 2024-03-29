package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bird struct {
	position        rl.Vector2
	currentFrame    uint8
	textures        [2]*rl.Texture2D
	lastTimeUpdate  float32
	frameTimeLength float32
}

func NewBird(position rl.Vector2) *Bird {
	texture1 := rl.LoadTexture("resources/bird/fly/frame-1.png")
	texture2 := rl.LoadTexture("resources/bird/fly/frame-2.png")
	//width := int32(float32(config.WindowWidth) * float32(0.1))
	//height := int32(float32(config.WindowHeight) * float32(0.15))
	return &Bird{
		position:        position,
		textures:        [2]*rl.Texture2D{&texture1, &texture2},
		frameTimeLength: float32(0.1),
	}
}

func (b *Bird) Draw() {
	if b.lastTimeUpdate > b.frameTimeLength {
		b.currentFrame += 1
		if b.currentFrame >= 2 {
			b.currentFrame = 0
		}
		b.lastTimeUpdate = 0
	} else {
		b.lastTimeUpdate += rl.GetFrameTime()
	}
	texture := b.textures[b.currentFrame]
	rl.DrawTextureEx(*texture, b.position, 0, 0.25, rl.White)
}

func (b *Bird) Update() {
	if rl.IsKeyDown(rl.KeySpace) {
		b.position.Y -= 10
	}
	b.position.Y += 2.5
}

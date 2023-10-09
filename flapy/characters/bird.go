package characters

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type bird struct {
	length       float32
	width        float32
	position     rl.Vector2
	texture      rl.Texture2D
	velocityLeft float32
}

type BirdParams struct {
	Length   float32
	Width    float32
	Position rl.Vector2
	Texture  rl.Texture2D
}

var (
	gravity  = 200
	velocity = 2000
)

func NewBird(p BirdParams) *bird {
	return &bird{
		length:   p.Length,
		width:    p.Width,
		position: p.Position,
		texture:  p.Texture,
	}
}

func (b bird) Length() float32 {
	return b.length
}

func (b bird) Width() float32 {
	return b.width
}

func (b bird) Position() rl.Vector2 {
	return b.position
}

func (b bird) Texture() rl.Texture2D {
	return b.texture
}

func (b bird) Draw() {
	sourceRec := rl.Rectangle{0, 0, float32(b.texture.Width), float32(b.texture.Height)}
	destRec := rl.Rectangle{b.position.X, b.position.Y, b.width, b.length}
	rl.DrawTexturePro(b.texture, sourceRec, destRec, rl.NewVector2(0, 0), 0, rl.White)
}

func (b *bird) Update() {
	delta := rl.GetFrameTime()
	if rl.IsKeyPressed(rl.KeySpace) {
		b.position.Y -= float32(velocity) * delta
		b.velocityLeft = float32(velocity)
	} else if b.velocityLeft > 0 {
		b.velocityLeft -= float32(gravity)
		b.position.Y -= b.velocityLeft * delta
	} else {
		b.position.Y += float32(gravity) * delta
	}
}

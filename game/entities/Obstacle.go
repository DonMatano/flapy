package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Obstacle struct {
	position rl.Vector2
	width    float32
	height   float32
	speed    float32
	isTop    bool
}

func NewObstacle(position rl.Vector2, width, height float32, isTop bool) *Obstacle {
	return &Obstacle{
		width:    width,
		height:   height,
		position: position,
		speed:    2.5,
		isTop:    isTop,
	}
}

func (o *Obstacle) Draw() {
	rotation := float32(0)
	height := o.height
	if !o.isTop {
		rotation = 180
		height = -o.height
	}

	//rl.DrawRectangle(int32(o.position.X), int32(o.position.Y), int32(o.width), int32(o.height), rl.Red)
	rl.DrawRectanglePro(
		rl.NewRectangle(o.position.X, o.position.Y, o.width, height),
		rl.NewVector2(0, 0),
		rotation,
		rl.Red,
	)
}

func (o *Obstacle) Update() {
	o.position.X -= o.speed
}

func (o *Obstacle) IsOutOfScreen() bool {
	return o.position.X+o.width < -50
}

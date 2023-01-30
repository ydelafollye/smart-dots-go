package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Goal struct {
	pos rl.Vector2
}

func newGoal() Goal {
	return Goal{rl.Vector2{screenWidth / 2, 10}}
}

func (g Goal) show() {
	rl.DrawCircle(int32(g.pos.X), int32(g.pos.Y), 10, rl.Color{255, 0, 0, 255})
}

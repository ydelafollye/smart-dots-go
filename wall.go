package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Wall struct {
	pos    rl.Vector2
	width  int32
	height int32
}

func newWall(pos rl.Vector2, width int32, height int32) Wall {
	return Wall{pos, width, height}
}

func (w Wall) show() {
	rl.DrawRectangle(int32(w.pos.X), int32(w.pos.Y), w.width, w.height, rl.Black)
}

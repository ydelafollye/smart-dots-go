package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

func limitVector2(v rl.Vector2, max float32) rl.Vector2 {
	if magSq(v) > math.Pow(float64(max), 2) {
		vLimit := rl.Vector2Normalize(v)
		mult(&vLimit, max)
		return vLimit
	}
	return v
}

func magSq(v rl.Vector2) float64 {
	return math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2)
}

func mult(v *rl.Vector2, n float32) {
	v.X = v.X * n
	v.Y = v.Y * n
}

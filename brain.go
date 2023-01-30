package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
)

type Brain struct {
	directions []rl.Vector2
	step       int
}

func newBrain() Brain {
	// rand.Seed(time.Now().UnixNano())
	var b Brain
	b.directions = make([]rl.Vector2, brainSize)
	b.step = 0
	b.randomize()
	return b
}

func (b Brain) clone() Brain {
	var newDirections []rl.Vector2
	newDirections = make([]rl.Vector2, brainSize)
	copy(newDirections, b.directions)
	return Brain{newDirections, 0}
}

func (b *Brain) randomize() {
	for i := range b.directions {
		randomAngle := rand.Float64() * 2 * math.Pi
		b.directions[i] = rl.Vector2{float32(math.Cos(randomAngle)), float32(math.Sin(randomAngle))}
	}
}

func (b *Brain) mutate() {
	mutationRate := 0.01
	for i := range b.directions {
		r := rand.Float64()
		if mutationRate > r {
			randomAngle := rand.Float64() * 2 * math.Pi
			b.directions[i] = rl.Vector2{float32(math.Cos(randomAngle)), float32(math.Sin(randomAngle))}
		}
	}
}

package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Dot struct {
	pos    rl.Vector2
	radius float32

	vel rl.Vector2
	acc rl.Vector2

	dead bool

	brain       Brain
	reachedGoal bool
	isBest      bool
	fitness     float64
}

func newDot() Dot {
	return Dot{rl.NewVector2(screenWidth/2, screenHeight), 5, rl.Vector2{0, 0}, rl.Vector2{0, 0}, false, newBrain(), false, false, 0}
}
func (d *Dot) move() {
	if len(d.brain.directions) > d.brain.step {
		d.acc = d.brain.directions[d.brain.step]
		d.brain.step++
	} else {
		d.dead = true
	}
	d.vel = rl.Vector2Add(d.vel, d.acc)
	d.vel = limitVector2(d.vel, 5)
	d.pos = rl.Vector2Add(d.pos, d.vel)
}

func (d *Dot) isDead(w []Wall) {
	// Hit a wall
	for i := range w {
		if rl.CheckCollisionCircleRec(d.pos, d.radius, rl.Rectangle{w[i].pos.X, w[i].pos.Y, float32(w[i].width), float32(w[i].height)}) {
			d.dead = true
		}
	}
	// Hit screen borders
	if d.pos.X <= 0 || d.pos.X >= screenWidth || d.pos.Y <= 0 || d.pos.Y >= screenHeight {
		d.dead = true
	}
}

func (d *Dot) hasReachedGoal(g Goal) {
	if rl.CheckCollisionCircles(d.pos, d.radius, g.pos, 10) {
		d.reachedGoal = true
	}

}
func (d Dot) show() {
	if d.isBest {
		rl.DrawCircle(int32(d.pos.X), int32(d.pos.Y), 6, rl.Color{0, 255, 0, 255})
	} else {
		rl.DrawCircle(int32(d.pos.X), int32(d.pos.Y), d.radius, rl.Black)
	}

}

func (d *Dot) update(g Goal, w []Wall) {
	if !d.dead && !d.reachedGoal {
		d.move()
		d.hasReachedGoal(g)
		d.isDead(w)
	}
}

func (d *Dot) calculateFitness(g Goal) {
	if d.reachedGoal {
		d.fitness = 1.0/16.0 + 10000.0/float64(d.brain.step*d.brain.step)
	} else {
		distanceToGoal := rl.Vector2Distance(d.pos, g.pos)
		d.fitness = float64(1.0 / (distanceToGoal * distanceToGoal))

	}
}

func (d Dot) makeBaby() Dot {
	baby := newDot()
	baby.brain = d.brain.clone()
	return baby
}

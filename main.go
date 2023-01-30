package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"strconv"
	"time"
)

const (
	screenWidth    = 600
	screenHeight   = 800
	populationSize = 400
	brainSize      = 400
)

type Game struct {
	WindowsShouldClose bool
	population         Population
	goal               Goal
	walls              []Wall
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.WindowsShouldClose = false
	g.walls = make([]Wall, 2)
	g.walls[0] = newWall(rl.Vector2{0, 300}, 400, 10)
	g.walls[1] = newWall(rl.Vector2{200, 550}, 400, 10)

	g.population = newPopulation()
	g.goal = newGoal()

}

func (g *Game) Update() {
	if rl.WindowShouldClose() {
		g.WindowsShouldClose = true
	}
	if g.population.gen == 0 {
		g.population.mutateBabies()
		g.population.gen++
	} else if g.population.allDotsDead() {
		g.population.calculateFitnesses(g.goal)
		g.population.naturalSelection(g.goal)
		g.population.mutateBabies()
	}
	g.population.update(g.goal, g.walls)
}

func (g Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	fps := strconv.Itoa(int(rl.GetFPS()))
	rl.DrawText("FPS:"+fps, 10, 10, 20, rl.Black)
	for i := range g.walls {
		g.walls[i].show()
	}
	g.goal.show()
	g.population.show()

	rl.EndDrawing()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := NewGame()
	rl.InitWindow(screenWidth, screenHeight, "Smart Dot")
	rl.SetTargetFPS(60)

	for !game.WindowsShouldClose {
		game.Update()
		game.Draw()
	}

	rl.CloseWindow()
}

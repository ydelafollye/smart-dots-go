package main

import (
	"math/rand"
)

type Population struct {
	dots       []Dot
	fitnessSum float64
	gen        int
	bestDot    Dot
	minStep    int
}

func newPopulation() Population {
	d := make([]Dot, populationSize)
	for i := range d {
		d[i] = newDot()
	}
	return Population{d, 0, 0, d[0], 1000}
}

func (p Population) show() {
	for i := range p.dots {
		p.dots[i].show()
	}
}

func (p *Population) update(g Goal, w []Wall) {
	for i := range p.dots {
		if p.dots[i].brain.step > p.minStep {
			p.dots[i].dead = true
		} else {
			p.dots[i].update(g, w)
		}
	}
}

func (p *Population) calculateFitnesses(g Goal) {
	for i := range p.dots {
		p.dots[i].calculateFitness(g)
	}
}

func (p *Population) calculateFitnessSum() {
	p.fitnessSum = 0.0
	for i := range p.dots {
		p.fitnessSum += p.dots[i].fitness
	}
}

func (p Population) allDotsDead() bool {
	for i := range p.dots {
		if !p.dots[i].dead && !p.dots[i].reachedGoal {
			return false
		}
	}
	return true
}

func (p *Population) mutateBabies() {
	for i := 1; i < len(p.dots); i++ {
		p.dots[i].brain.mutate()
	}
}
func (p *Population) setBestDot() {
	for i := range p.dots {
		if p.dots[i].fitness > p.bestDot.fitness {
			p.bestDot = p.dots[i]
		}
	}
	if p.bestDot.reachedGoal {
		p.minStep = p.bestDot.brain.step
	}
}

func (p *Population) selectMommy() Dot {
	r := rand.Float64() * p.fitnessSum
	runningSum := 0.0
	for i := range p.dots {
		runningSum += p.dots[i].fitness
		if runningSum > r {
			return p.dots[i]
		}
	}
	return newDot()
}

func (p *Population) naturalSelection(g Goal) {
	var newPop []Dot
	newPop = make([]Dot, populationSize)
	p.setBestDot()
	p.calculateFitnessSum()

	newPop[0] = p.bestDot.makeBaby()
	newPop[0].isBest = true
	for i := 1; i < len(newPop); i++ {
		mommy := p.selectMommy()
		newPop[i] = mommy.makeBaby()
	}
	p.dots = newPop
	p.gen++
}

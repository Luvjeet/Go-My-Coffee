package particles

import (
	"math"
	"time"
)

type Particle struct {
	lifetime int
	speed    float64

	x float64
	y float64
}

type NextPosition func(particle *Particle, deltaMS int64)
type Ascii func(x, y int, count [][]int) rune

type ParticleParams struct {
	Maxlife  int64
	MaxSpeed float64

	ParticleCount int

	X int
	Y int

	nextPosition NextPosition
	ascii        Ascii
}

type ParticleSystem struct {
	ParticleParams
	particles []*Particle

	lastTime int64
}

func NewParticleSystem(params ParticleParams) ParticleSystem {
	return ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
	}
}

func (ps *ParticleSystem) Update() {
	now := time.Now().UnixMilli()
	delta := now - ps.lastTime
	ps.lastTime = now

	for _, p := range ps.particles {
		ps.nextPosition(p, delta)
	}
}

func (ps *ParticleSystem) Display() [][]rune {
	counts := make([][]int, 0)
	for row := 0; row < ps.Y; row++ {
		count := make([]int, 0)
		for col := 0; col < ps.X; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}

	for _, p := range ps.particles {
		row := int(math.Floor(p.y))
		col := int(math.Floor(p.x))
	}
}

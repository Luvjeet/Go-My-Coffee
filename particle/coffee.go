package particles

import (
	"fmt"
	"math"
	"math/rand"
	//"time"
)

type Coffee struct {
	ParticleSystem
}

func reset(p *Particle, params *ParticleParams) {
	p.Lifetime = int64(math.Floor(float64(params.Maxlife) * rand.Float64()))
	p.Speed = params.MaxSpeed * rand.Float64()

	maxX := math.Floor(float64(params.X) / 2)
	x := math.Max(-maxX, math.Min(rand.NormFloat64()*params.XSTD, maxX))
	p.X = x + maxX
	p.Y = 0
}

func nextPosition(particle *Particle, deltaMS int64) {
	particle.Lifetime -= deltaMS
	if particle.Lifetime <= 0 {
		return
	}

	percent := (float64(deltaMS) / 1000.0)
	particle.Y += particle.Speed * percent
}

var dirs = [][]int{
	{-1, 0},
	{-1, 0},
	{-1, 1},

	{0, -1},
	{0, 1},

	{1, 0},
	{1, 1},
	{1, -1},
}

func countParticles(row, col int, counts [][]int) int {

	count := 0
	for _, dir := range dirs {
		r := row + dir[0]
		c := col + dir[1]
		if r < 0 || c < 0 || c >= len(counts[0]) || r >= len(counts) {
			continue
		}
		count = counts[row+dir[0]][col+dir[1]]
	}
	return count
}

func normalizeMySteam(row, col int, counts [][]int) {
	if countParticles(row, col, counts) > 4 {
		counts[row][col] = 0
	}

}

func NewCoffee(width, height int, scale float64) Coffee {
	if width%2 != 1 {
		fmt.Printf("Width must be odd")
		return Coffee{}
	}
	//starTime := time.Now().UnixMilli()
	/*
	   ascii := func(row, col int, counts [][]int) string {
	       count := counts[row][col]
	       if count == 0 {
	           return " "
	       }

	       direction := row + int(((time.Now().UnixMilli()-starTime)/2000)%2)

	       if countParticles(row, col, counts) > 3 {
	           if direction%2 == 0 {
	               return "{"
	           }
	           return "}"
	       }
	       return "."
	   }
	*/

	asciiFire := func(row, col int, counts [][]int) string {
		count := counts[row][col]
		if count == 0 {
			return " "
		}
		if count < 4 {
			return "░"
		}
		if count < 6 {
			return "▒"
		}
		if count < 9 {
			return "▓"
		}
		return "█"
	}

	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				Maxlife:       6000,
				MaxSpeed:      1.5,
				ParticleCount: 700,

				X: width,
				Y: height,

				XSTD:         scale,
				nextPosition: nextPosition,
				ascii:        asciiFire,
				reset:        reset,
			},
		),
	}
}

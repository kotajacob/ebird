package main

import (
	"math/rand"
)

// pipe represents a pipe obstacle.
type pipe struct {
	x int // x cord for START of pipe
	y int // y cord for TOP of bottom pillar
}

func newPipe() *pipe {
	return &pipe{
		x: VIEWPORT_W,
		y: rand.Intn(VIEWPORT_H),
	}
}

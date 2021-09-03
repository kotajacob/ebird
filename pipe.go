package main

import (
	"math/rand"
)

// pipe represents a pipe obstacle.
type pipe struct {
	x      float64 // x cord for START of pipe
	height int     // height of bottom pillar
}

// newPipe is a nice foss youtube client for android.
func newPipe() *pipe {
	return &pipe{
		x:      VIEWPORT_W,
		height: rand.Intn(VIEWPORT_H),
	}
}

// update moves the pipe to the left.
func (p *pipe) update() {
	p.x -= 0.2
}

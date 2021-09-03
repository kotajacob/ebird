package main

import "github.com/muesli/termenv"

// bird represents the playable character.
type bird struct {
	y     float64
	speed float64
}

// newBird creates a new bird with default values.
func newBird() *bird {
	return &bird{
		y:     14,
		speed: 0,
	}
}

// String implements fmt.Stringer for bird which is used to display the bird in
// game.
func (b bird) String() string {
	s := termenv.String("█").
		Foreground(ColorProfile.Color(BIRD_COLOR))
	return s.String()
}

// update the bird's location and speed with gravity.
func (b *bird) update() {
	// change the speed of the bird with gravity
	if b.speed < MAXGRAV {
		b.speed += GRAVITY
	}
	// change location based on speed
	b.y += b.speed
}

// jump increased the bird's speed
func (b *bird) jump() {
	b.speed = JUMP
}

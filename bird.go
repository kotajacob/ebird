package main

// bird represents the playable character.
type bird struct {
	x     float64
	y     float64
	speed float64
}

// newBird creates a new bird with default values.
func newBird() *bird {
	return &bird{
		x:     20,
		y:     14,
		speed: 0,
	}
}

// update the bird's location and speed with gravity.
func (b *bird) update() {
	// change the speed of the bird with gravity
	if b.speed > -MAXSPEED {
		b.speed += GRAVITY
	}
	// change location based on speed
	b.y += b.speed
}

// jump increased the bird's speed
func (b *bird) jump() {
	b.speed = JUMP
}

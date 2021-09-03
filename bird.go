package main

import "github.com/charmbracelet/lipgloss"

// bird represents the playable character.
type bird struct {
	y     float64
	speed float64
	style lipgloss.Style
}

// newBird creates a new bird with default values.
func newBird() *bird {
	return &bird{
		y:     14,
		speed: 0,
		style: lipgloss.NewStyle().Foreground(lipgloss.Color(BIRD_COLOR)),
	}
}

// String implements fmt.Stringer for bird which is used to display the bird in
// game.
func (b bird) String() string {
	return b.style.Render("█")
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

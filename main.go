package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// 0,0 is in the top left and the board is drawn from the top left.
// As a result gravity is positive and jump is negative.
const (
	INTERVAL   = time.Second / 60 // tick rate
	VIEWPORT_W = 80               // viewport width
	VIEWPORT_H = 24               // viewport height
	GRAVITY    = 0.01             // gravity constant
	MAXGRAV    = 0.32             // maximum gravity
	JUMP       = -0.4             // jump speed
	BIRD_X     = 18               // bird x cordinate
	BIRD_COLOR = "11"             // ANSI color for the bird
)

// model is a tea.Model representing the ebird game.
type model struct {
	bird *bird
}

// newModel creates a new model with default values.
func newModel() model {
	b := newBird()
	return model{
		bird: b,
	}
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k", "w", "enter", " ":
			m.bird.jump()
		}
	case tickMsg:
		m.bird.update()
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
	var s strings.Builder
	by := int(math.Round(m.bird.y))
	for y := 0; y < VIEWPORT_H-1; y++ {
		s.WriteString("│")
		for x := 0; x < VIEWPORT_W-2; x++ {
			switch x {
			case BIRD_X:
				if y == by {
					s.WriteString(m.bird.String())
				} else {
					s.WriteString(" ")
				}
			default:
				s.WriteString(" ")
			}
		}
		s.WriteString("│\n")
	}
	return s.String()
}

func main() {
	m := newModel()
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("game broke :(", err)
		os.Exit(1)
	}
}

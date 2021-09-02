package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	VIEWPORT_W = 80               // viewport width
	VIEWPORT_H = 24               // viewport height
	GRAVITY    = -0.005           // gravity constant
	MAXSPEED   = 0.3              // maximum speed
	JUMP       = 0.3              // jump speed
	INTERVAL   = time.Second / 60 // tick rate
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
	return fmt.Sprintf("bird y: %v\nbird speed: %v\n", m.bird.y, m.bird.speed)
}

func main() {
	m := newModel()
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("game broke :(", err)
		os.Exit(1)
	}
}

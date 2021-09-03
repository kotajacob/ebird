package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

// 0,0 is in the top left and the board is drawn from the top left.
// As a result gravity is positive and jump is negative. Physics is tied to FPS
// cause I was lazy, perhaps there's an easy way to use delta time I could
// investigate.
const (
	INTERVAL   = time.Second / 60 // tick rate
	VIEWPORT_W = 80               // viewport width
	VIEWPORT_H = 24               // viewport height
	SPEED      = 1 / 60           // speed of incoming pipes per tick
	GRAVITY    = 0.01             // gravity rate per tick
	MAXGRAV    = 0.32             // maximum gravity per tick
	JUMP       = -0.4             // jump speed
	BIRD_X     = 18               // bird x cordinate
	PIPE_W     = 3                // pipe width
	PIPE_GAP   = 5                // pipe gap
)

var (
	ColorProfile = termenv.ColorProfile()

	BirdString = termenv.String("█").
			Foreground(ColorProfile.Color("11"))
	PipeString = termenv.String("█").
			Foreground(ColorProfile.Color("10")).
			Background(ColorProfile.Color("10"))
)

// model is a tea.Model representing the ebird game.
type model struct {
	bird  *bird
	pipes []*pipe
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
		for _, p := range m.pipes {
			p.update()
		}
		m.bird.update()
		return m, tick()
	}
	return m, nil
}

// getXY returns a character for an x,y cordinate.
func (m model) getXY(x, y int) string {
	// render sky
	s := " "
	// check pipes
	for _, p := range m.pipes {
		xx := int(math.Round(p.x))
		if x >= xx && x <= xx+PIPE_W {
			if y <= p.height || y > p.height+PIPE_GAP {
				s = PipeString.String()
			}
		}
	}
	// check player
	if x == BIRD_X {
		if y == int(math.Round(m.bird.y)) {
			s = BirdString.String()
		}
	}
	return s
}

func (m model) View() string {
	var s strings.Builder
	for y := 0; y < VIEWPORT_H; y++ {
		for x := 0; x < VIEWPORT_W; x++ {
			s.WriteString(m.getXY(x, y))
		}
		if y < VIEWPORT_H-1 { // skip last newline
			s.WriteString("\n")
		}
	}
	return s.String()
}

func main() {
	rand.Seed(time.Now().Unix())
	m := newModel()
	p := newPipe()
	m.pipes = append(m.pipes, p)
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("game broke :(", err)
		os.Exit(1)
	}
}

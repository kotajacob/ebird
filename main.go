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
	FREQ       = time.Second * 2  // frequency of new pipes
	SPEED      = 0.2              // speed of incoming pipes per tick
	GRAVITY    = 0.016            // gravity rate per tick
	MAXGRAV    = 0.34             // maximum gravity per tick
	JUMP       = -0.38            // jump speed
	BIRD_X     = 18               // bird x cordinate
	PIPE_W     = 6                // pipe width
	PIPE_GAP   = 8                // pipe gap
)

var (
	ColorProfile = termenv.ColorProfile()

	BirdString = termenv.String("█").
			Foreground(ColorProfile.Color("11")).
			Background(ColorProfile.Color("11"))
	PipeString = termenv.String("█").
			Foreground(ColorProfile.Color("10")).
			Background(ColorProfile.Color("10"))
)

// model is a tea.Model representing the ebird game.
type model struct {
	bird     *bird
	pipes    []*pipe
	lastPipe time.Time
}

// newModel creates a new model with default values.
func newModel() model {
	b := newBird()
	l := time.Now()
	return model{
		bird:     b,
		lastPipe: l,
	}
}

func (m model) Init() tea.Cmd {
	return tick()
}

func collision(b *bird, p *pipe) bool {
	if b.y > VIEWPORT_H {
		return true
	}
	// check pipe
	px := int(math.Round(p.x))
	by := int(math.Round(b.y))
	if BIRD_X >= px && BIRD_X <= px+PIPE_W {
		if by <= p.height || by > p.height+PIPE_GAP {
			return true
		}
	}
	return false
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
		// spawn new pipe if needed
		t := time.Time(msg)
		if t.After(m.lastPipe.Add(FREQ)) {
			np := newPipe()
			m.pipes = append(m.pipes, np)
			m.lastPipe = t
		}
		// update the bird
		m.bird.update()
		// move pipes and check for collisions
		for _, p := range m.pipes {
			p.update()
			if collision(m.bird, p) {
				return m, tea.Quit
			}
		}
		return m, tick()
	}
	return m, nil
}

// getXY returns a character for an x,y coordinate.
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
	if x >= BIRD_X && x <= BIRD_X+2 {
		yy := int(math.Round(m.bird.y))
		if y >= yy && y <= yy+1 {
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
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("game broke :(", err)
		os.Exit(1)
	}
}

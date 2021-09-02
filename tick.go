package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Duration(INTERVAL), func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

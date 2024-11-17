package app

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Root struct {
	*tea.Program
}

func New() *Root {
	p := tea.NewProgram(newApp(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	return &Root{
		p,
	}
}

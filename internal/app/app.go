package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	*tea.Program
}

func New() *App {
	return &App{}
}

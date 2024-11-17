package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tommjj/tasks/internal/app/ui/taskinput"
	"github.com/tommjj/tasks/internal/core/ports"
)

type (
	errMsg error
)

type App struct {
	Repo ports.ITaskRepository

	textInput *taskinput.Model
	err       error
}

func newApp() *App {
	ti := taskinput.New()

	return &App{
		textInput: ti,
		err:       nil,
	}
}

func (m *App) Init() tea.Cmd {
	return textinput.Blink
}

func (m *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.textInput.SetValue("")
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m *App) View() string {
	return fmt.Sprintf(
		"Add task\n%s%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

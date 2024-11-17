package taskinput

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tommjj/tasks/internal/core/domain"
)

var priorityIcon = map[domain.Priority]string{
	domain.Pri1: "🚩",
	domain.Pri2: "⏫",
	domain.Pri3: "🔼",
	domain.Pri4: ">>",
}

var (
	priorityUpKey   = tea.KeyCtrlP
	priorityDownKey = tea.KeyCtrlO
)

type Model struct {
	textinput.Model

	priority domain.Priority
}

func New() *Model {
	input := textinput.New()
	input.Placeholder = "new task"
	input.Focus()
	input.CharLimit = 255
	input.Width = 20
	input.Prompt = " "

	return &Model{
		Model:    input,
		priority: domain.Pri4,
	}
}

func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *Model) GetPriority() domain.Priority {
	return m.priority
}

func (m *Model) SetPriority(priority domain.Priority) {
	m.priority = priority
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// handle set priority
	case tea.KeyMsg:
		switch msg.Type {
		case priorityUpKey:
			m.priorityUp()
			return m, nil
		case priorityDownKey:
			m.priorityDown()
			return m, nil
		}

	// handle set width when window size change
	case tea.WindowSizeMsg:
		m.Width = msg.Width - 7
		return m, nil
	}

	m.Model, cmd = m.Model.Update(msg)

	return m, cmd
}

func (m *Model) View() string {
	return fmt.Sprintf(
		"%s %s", priorityIcon[m.priority],
		m.Model.View(),
	) + "\n"
}

func (m *Model) priorityUp() {
	if m.priority == domain.Pri1 {
		m.priority = domain.Pri4
	} else {
		m.priority += 1
	}
}

func (m *Model) priorityDown() {
	if m.priority == domain.Pri4 {
		m.priority = domain.Pri1
	} else {
		m.priority -= 1
	}
}

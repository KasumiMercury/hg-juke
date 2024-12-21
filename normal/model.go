package normal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width, height int
}

func New(width, height int) tea.Model {
	return &model{width, height}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Height(m.height).
		Width(m.width).
		Render("content")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

package setting

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width, height int
}

func New() Model {
	return Model{}
}

func (m Model) Build(_ string, width, height int) (tea.Model, error) {
	return &Model{width, height}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		Height(m.height).
		Width(m.width).
		Align(lipgloss.Center).
		Render("setting")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

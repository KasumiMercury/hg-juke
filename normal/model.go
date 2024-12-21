package normal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strconv"
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
		Height(m.height).
		Width(m.width).
		Align(lipgloss.Center).
		Render(strconv.Itoa(m.width))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

package top

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width  int
	height int
}

func newModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m model) View() string {
	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(1).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		Render("header")

	footer := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(1).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		Render("footer")

	content := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(m.height - lipgloss.Height(header) - lipgloss.Height(footer)).
		Render()

	display := lipgloss.JoinVertical(lipgloss.Top, header, content, footer)

	return display
}

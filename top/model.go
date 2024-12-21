package top

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"hg-juke/normal"
	"hg-juke/page"
)

type model struct {
	current page.Page
	width   int
	height  int
}

func newModel() model {
	nm := normal.New(100, 5)
	p := page.New("normal", nm)
	return model{
		current: *p,
	}
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

	content := m.current.View()

	display := lipgloss.JoinVertical(lipgloss.Top, header, content, footer)

	return display
}

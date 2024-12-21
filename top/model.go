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
	m := model{}

	nm := normal.New(m.width, m.height)
	p := page.New("normal", nm)
	m.current = *p
	return m
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
		m.recalculateContentSize()
	}

	return m, nil
}

func (m model) View() string {
	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(1).
		Render("header")

	footer := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(1).
		Render("footer")

	content := lipgloss.NewStyle().
		Height(m.contentHeight()).
		Width(m.contentWidth()).
		Render(m.current.View())

	display := lipgloss.JoinVertical(lipgloss.Top, header, content, footer)

	return display
}

func (m model) recalculateContentSize() {
	w := m.contentWidth()
	h := m.contentHeight()
	_ = m.current.Update(tea.WindowSizeMsg{
		Width:  w,
		Height: h,
	})
}

func (m model) contentWidth() int {
	return m.width
}
func (m model) contentHeight() int {
	return m.height - 2
}

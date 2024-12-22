package top

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"hg-juke/normal"
	"hg-juke/page"
	"hg-juke/router"
	"hg-juke/setting"
	"log"
)

type model struct {
	*router.Router
	width, height int
}

func newTop(isInitial bool) model {
	m := model{
		router.NewRouter(),
		0,
		0,
	}

	m.SetBuilder(page.Normal, normal.Model{})
	m.SetBuilder(page.Setting, setting.Model{})

	first := page.Normal
	title := ""

	if isInitial {
		first = page.Setting
		title = "initial"
	}

	err := m.SetPage(first, "", title)
	if err != nil {
		log.Fatal(err)
		return model{}
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		default:
			cmd := m.Current.Update(msg)
			return m, cmd
		}
	case router.NewPageMsg:
		if err := m.SetPage(msg.PageType, msg.Id, msg.Title); err != nil {
			// TODO: handling error
			log.Fatal(err)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.recalculateContentSize()
	default:
		cmd := m.Current.Update(msg)
		return m, cmd
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
		Render(m.Current.View())

	display := lipgloss.JoinVertical(lipgloss.Top, header, content, footer)

	return display
}

func (m model) recalculateContentSize() {
	w := m.contentWidth()
	h := m.contentHeight()
	m.ChangeSize(w, h)

	_ = m.Current.Update(tea.WindowSizeMsg{
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

package setting

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"hg-juke/config"
	"hg-juke/page"
	"hg-juke/router"
	"log"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#1ED760"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle
	noStyle      = lipgloss.NewStyle()

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type Model struct {
	width, height int
	focusIndex    int
	inputs        []textinput.Model
	miss          bool
}

func (m Model) Build(_ string, width, height int) (tea.Model, error) {
	m.width = width
	m.height = height

	m.inputs = make([]textinput.Model, 3)
	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Spotify Client ID"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Spotify Client Secret"
		case 2:
			t.Placeholder = "Misskey API Key"
		}

		m.inputs[i] = t
	}

	return &m, nil
}

type missField struct{}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *Model) View() string {
	elms := make([]string, len(m.inputs))

	for i := range m.inputs {
		elm := lipgloss.NewStyle().Padding(1, 2).Render(m.inputs[i].View())
		elms[i] = elm
	}

	inputs := lipgloss.JoinVertical(lipgloss.Left, elms...)

	cmps := make([]string, 0, 3)
	if m.miss {
		cmps = append(cmps, lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("All inputs must be filled"))
	}

	cmps = append(cmps, inputs)

	button := blurredButton
	if m.focusIndex == len(m.inputs) {
		button = focusedButton
	}
	cmps = append(cmps, lipgloss.NewStyle().Padding(1, 2).Render(button))
	panel := lipgloss.JoinVertical(lipgloss.Center, cmps...)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, panel)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case missField:
		m.miss = true
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, m.submit
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}

				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m Model) submit() tea.Msg {
	for i, input := range m.inputs {
		v := input.Value()
		if v == "" {
			return missField{}
		}

		switch i {
		case 0:
			config.Set("SpotifyClientID", v)
		case 1:
			config.Set("SpotifyClientSecret", v)
		case 2:
			config.Set("MisskeyAPIKey", v)
		}
	}

	if err := config.Write(); err != nil {
		//TODO: handling error
		log.Fatal(err)
	}

	// TODO: exit setting page

	msg := router.NewPageMsg{
		PageType: page.Normal,
		Title:    "top",
	}
	return msg
}

package top

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func Start() {
	m := newModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

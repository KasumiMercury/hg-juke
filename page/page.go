package page

import tea "github.com/charmbracelet/bubbletea"

type Page struct {
	title string
	model tea.Model
}

func New(title string, model tea.Model) *Page {
	return &Page{
		title: title,
		model: model,
	}
}

func (p Page) Update(msg tea.Msg) tea.Cmd {
	m, cmd := p.model.Update(msg)
	p.model = m
	return cmd
}

func (p Page) View() string {
	return p.model.View()
}

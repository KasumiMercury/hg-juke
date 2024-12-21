package top

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"hg-juke/page"
)

type Router struct {
	Current       *page.Page
	builder       map[page.Type]Builder
	width, height int
}

type Builder interface {
	Build(id string, width, height int) (tea.Model, error)
}

func NewRouter() *Router {
	return &Router{
		builder: make(map[page.Type]Builder),
	}
}

func (r *Router) SetBuilder(pType page.Type, builder Builder) {
	r.builder[pType] = builder
}

func (r *Router) SetPage(pType page.Type, id string, title string) error {
	b, ok := r.builder[pType]
	if !ok {
		return fmt.Errorf("no builder found for type %s", pType)
	}

	m, err := b.Build(id, r.width, r.height)
	if err != nil {
		return err
	}

	r.Current = page.New(title, m)

	return nil
}

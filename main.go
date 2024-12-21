package main

import (
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"hg-juke/config"
	"log"
	"log/slog"
	"os"
)

type model struct {
	width, height int
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

func main() {
	logLevel := new(slog.LevelVar)
	logOpts := slog.HandlerOptions{
		Level: logLevel,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &logOpts))
	slog.SetDefault(logger)

	isDebug := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	if isDebug != nil && *isDebug {
		logLevel.Set(slog.LevelDebug)
	}

	confExist, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	if !confExist {
		fmt.Println("config file not exists, create config file")
		// TODO: initialize sequence
		return
	}

	p := tea.NewProgram(model{}, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

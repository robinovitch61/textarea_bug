package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"tmp/textarea"
)

type model struct {
	textarea textarea.Model
	ready    bool
	err      error
}

func main() {
	if err := start(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func start() error {
	m := model{
		textarea: textarea.New(),
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	return p.Start()
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		if !m.ready {
			m.textarea.View()
			panic("here")
			m.ready = true
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	return m.textarea.View()
}

package main

import (
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	keyboardLayout   string
	keyboardSize     int
	isInfoBarVisible bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+shift+h":
			m.isInfoBarVisible = !m.isInfoBarVisible
			return m, nil
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	s := getView(m)
	if !m.isInfoBarVisible {
		s = ""
	}
	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

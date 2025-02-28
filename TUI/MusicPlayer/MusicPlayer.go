package MusicPlayer

import (
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	// Add fields as needed
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// switch msg := msg.(type) {
	// // Handle updates
	// }
	return m, nil
}

func (m Model) View() string {
	return lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render("Music Player")
}

func New() Model {
	return Model{}
}

package TUI

import (
	Init "Melodex/TUI/Init"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
	"strings"
)

// Model struct (exported for use in MusicList)
type Model struct {
	List           table.Model
	TotalListWidth int
	width          int
	height         int
}

// Update handles input events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.List.MoveUp(1)
		case "down":
			m.List.MoveDown(1)
		case "pgup":
			m.List.MoveUp(10)
		case "pgdown":
			m.List.MoveDown(10)
		case "home":
			m.List.GotoTop()
		case "end":
			m.List.GotoBottom()
		case "enter":
			fmt.Printf("Selected: %s\n", m.List.SelectedRow()[0])
		}

	case tea.MouseMsg:
		switch msg.Button {
		case tea.MouseButtonLeft:
			if msg.Y >= 5 && msg.Y < (5+m.List.Height()) {
				rowIdx := msg.Y - 5
				m.List.SetCursor(rowIdx)
			}
		case tea.MouseButtonWheelUp:
			m.List.MoveUp(1)
		case tea.MouseButtonWheelDown:
			m.List.MoveDown(1)
		}

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.List.SetHeight(m.height - 5)
	}

	return m, cmd
}

// View renders the TUI
func (m Model) View() string {
	playlistName := "PLAYLISTNAME"
	staticText := " 󰉹 SEARCH  "

	padding := m.TotalListWidth - lipg.Width(playlistName) - lipg.Width(staticText) + 5
	if padding < 0 {
		padding = 0
	}

	header := lipg.NewStyle().BorderStyle(lipg.RoundedBorder()).Render(
		playlistName + strings.Repeat(" ", padding) + staticText)

	return header + "\n" + lipg.NewStyle().BorderStyle(lipg.RoundedBorder()).Render(m.List.View())
}

// Application runs the TUI
func Application() {

}

// Init initializes the TUI
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Melodex"),
	)
}

package MusicList

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	bzone "github.com/lrstanley/bubblezone"
)

// Init runs the TUI
func Init() {
	bzone.NewGlobal()
	p := tea.NewProgram(
		MList(Model{}), // This now correctly returns a tea.Model
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

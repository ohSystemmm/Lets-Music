package Init

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	bzone "github.com/lrstanley/bubblezone"
	"os"
)

// Init runs the TUI
func Init() {
	bzone.NewGlobal()
	p := tea.NewProgram(
		MList.MusicList(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

package TUI

import (
	TUI "Melodex/TUI/MusicList"
	"github.com/charmbracelet/bubbles/table"
)

type model struct {
	list           table.Model
	totalListWidth int
	width          int
	height         int
}

func Application() {
	TUI.Run()
}

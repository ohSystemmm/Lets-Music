package TUI

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	bzone "github.com/lrstanley/bubblezone"
)

type model struct {
	table  table.Model
	width  int
	height int
}

func List() model {
	columns := []table.Column{
		{Title: "Title", Width: 40},
		// {Title: "Category", Width: 8},
		// {Title: "Rating", Width: 8},
		// {Title: "Started", Width: 14},
		// {Title: "Finished", Width: 14},
		// {Title: "Status", Width: 12},
		// {Title: "Notes", Width: 40},
	}

	// rows := []table.Row{
	// 	{"1 2", "Anime", "100/99", "Sometime", "Sometime", "Finished", "Was good"},
	// 	{"1 2 Entry", "Manga", "90/100", "Yesterday", "Today", "In Progress", "Interesting!"},
	// 	{"1 2 Bad", "TV Show", "98/100", "Last Year", "Last Week", "Finished", "Amazing plot!"},
	// 	{"1 2", "Anime", "85/100", "2010", "2015", "Finished", "Classic, but long."},
	// 	{"1 2 Piece", "Anime", "99/100", "2000", "Ongoing", "Ongoing", "The adventure never ends!"},
	// 	{"1 2 Note", "Anime", "95/100", "2015", "2015", "Finished", "Thrilling!"},
	// 	{"1 2 on Titan", "Anime", "100/100", "2017", "2023", "Finished", "Unbelievable twists!"},
	// 	{"1 2 Name", "Movie", "92/100", "2018", "2018", "Finished", "Beautiful and emotional."},
	// 	{"1 2 Slayer", "Anime", "95/100", "2019", "Ongoing", "Ongoing", "Incredible animation."},
	// }
	rows := []table.Row{
		{"ASD"},
		{"asdwqr"},
		{"rfew"},
		{"fgdgs"},
		{"qwdsa"},
		{"frsd"},
		{"h56a"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(32),
	)

	t.SetStyles(defineTableStyles())

	return model{table: t}
}

func defineTableStyles() table.Styles {
	styles := table.DefaultStyles()
	styles.Selected = styles.Selected.
		Foreground(lipgloss.Color("230")).
		Background(lipgloss.Color("63")).
		Bold(true)
	styles.Header = styles.Header.Bold(true).Background(lipgloss.Color("60"))
	return styles
}

func Run() {
	bzone.NewGlobal()
	p := tea.NewProgram(
		List(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Ani-Track"),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.table.MoveUp(1)
		case "down":
			m.table.MoveDown(1)
		}

	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			if msg.Y >= 2 && msg.Y < (2+m.table.Height()) {
				rowIdx := msg.Y - 2
				m.table.SetCursor(rowIdx)
			}
		}

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.table.SetHeight((m.height / 2) - 3)
	}

	return m, cmd
}

func (m model) View() string {
	return lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Render(m.table.View())
}

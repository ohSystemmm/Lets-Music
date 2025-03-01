package MusicList

import (
	"Melodex/TUI/SharedState"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	SharedState *SharedState.SharedState

	List      table.Model
	SearchBar textinput.Model

	PlaylistName   string
	TotalListWidth int

	Width  int
	Height int
}

// Init implements the tea.Model interface
func (m Model) Init() tea.Cmd {
	return nil // No initial command needed
}

// Update handles input events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
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
		}

		if m.SharedState.Searching {
			switch msg.String() {
			case "esc", "enter":
				m.SharedState.Searching = false
				m.SearchBar.Blur()
				// return m, nil
			}
		} else {
			switch msg.String() {
			case "enter":
				// Handle selection
			case "f":
				m.SharedState.Searching = true
				return m, m.SearchBar.Focus()
			case "q":
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		// This is to handle the window resize
		m.Width, m.Height = msg.Width, msg.Height
		m.List.SetHeight(m.Height - 5)
		newColumns := m.List.Columns()

		// availableWidth := m.width - 10
		// titleWidth := int(0.3 * float64(availableWidth))
		// lengthWidth := int(0.2 * float64(titleWidth))

		// newColumns[0].Width = titleWidth
		// newColumns[1].Width = lengthWidth

		// m.TotalListWidth = titleWidth + lengthWidth
		m.List.SetColumns(newColumns)

	}
	m.SearchBar, cmd = m.SearchBar.Update(msg)

	return m, cmd
}

// View renders the music list
func (m Model) View() string {
	// FIX Text input fileed from bubbles for search
	searchBar := m.SearchBar.View()
	// filtering := " 󰉹 "
	filtering := " "
	// filtering := " "

	// filtering := " "
	// filtering := " "

	// filtering := "󱕉 "
	// filtering := "󱕋 "
	// filtering := "󱕊 "
	// filtering := "󱕌 "

	padding := m.TotalListWidth - lipg.Width(m.PlaylistName) - lipg.Width(filtering) - lipg.Width(searchBar) + 4
	padding = max(padding, 0)

	header := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(
		lipg.NewStyle().Bold(true).Render(m.PlaylistName) + strings.Repeat(" ", padding) + filtering + searchBar)

	musicList := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(m.List.View())

	return lipg.JoinVertical(lipg.Top, header, musicList)
}

// New initializes the music list
func New(sharedState *SharedState.SharedState) Model {
	rows := []table.Row{
		{"Song A", "3:40"},
		{"Song B", "4:20"},
		{"Song C", "2:50"},
		{"Song A", "3:40"},
		{"Song B", "4:20"},
		{"Song C", "2:50"},
	}

	longestTitle, longestTime := 98, 15

	for _, row := range rows {
		longestTitle = max(longestTitle, len(row[0]))
		longestTime = max(longestTime, len(row[1]))
	}

	columns := []table.Column{
		{Title: "Title", Width: longestTitle},
		{Title: "Length", Width: longestTime},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	t.SetStyles(defineTableStyles())

	tLW := 0
	for _, col := range columns {
		tLW += col.Width
	}

	sB := textinput.New()
	sB.Width = 20
	sB.CharLimit = 0
	sB.Placeholder = "Search"

	// FIX find a way to somehow seperate the filter and search boxes in their own borders
	sB.Prompt = " "

	return Model{
		SharedState:    sharedState,
		List:           t,
		TotalListWidth: tLW,
		PlaylistName:   "Example Playlistname",
		SearchBar:      sB,
	}
}

// defineTableStyles sets the table styles
func defineTableStyles() table.Styles {
	styles := table.DefaultStyles()
	styles.Selected = styles.Selected.
		Foreground(lipg.Color("230")).
		Background(lipg.Color("63")).
		Bold(true)
	styles.Header = styles.Header.Bold(true).Background(lipg.Color("60"))
	return styles
}

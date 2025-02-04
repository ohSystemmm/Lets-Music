package MusicList

import (
	"github.com/charmbracelet/bubbles/table"
	lipg "github.com/charmbracelet/lipgloss"
)

// MusicList initializes the music list
func MusicList() Model {
	rows := []table.Row{
		{"Song A", "3:40"},
		{"Song B", "4:20"},
		{"Song C", "2:50"},
	}

	longestTitle, longestTime := 13, 6
	for _, row := range rows {
		if len(row[0]) > longestTitle {
			longestTitle = len(row[0])
		}
		if len(row[1]) > longestTime {
			longestTime = len(row[1])
		}
	}

	columns := []table.Column{
		{Title: "Title", Width: longestTitle},
		{Title: "Length", Width: longestTime},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)
	t.SetStyles(defineTableStyles())

	tLW := 0
	for _, col := range columns {
		tLW += col.Width
	}

	return Model{List: t, TotalListWidth: tLW}
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

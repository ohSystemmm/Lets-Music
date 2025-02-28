package TUI

import (
	"Melodex/TUI/MusicList"
	"Melodex/TUI/MusicPlayer"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
	bzone "github.com/lrstanley/bubblezone"
	"os"
)

type MainModel struct {
	MList   MusicList.Model
	MPlayer MusicPlayer.Model
	width   int
	height  int
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.MList.Init(), m.MPlayer.Init())
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height

		var mList tea.Model
		mList, cmd = m.MList.Update(msg)
		m.MList = mList.(MusicList.Model)
		cmds = append(cmds, cmd)

		var mPlayer tea.Model
		mPlayer, cmd = m.MPlayer.Update(msg)
		m.MPlayer = mPlayer.(MusicPlayer.Model)
		cmds = append(cmds, cmd)
	default:
		var mList tea.Model
		mList, cmd = m.MList.Update(msg)
		m.MList = mList.(MusicList.Model)
		cmds = append(cmds, cmd)

		var mPlayer tea.Model
		mPlayer, cmd = m.MPlayer.Update(msg)
		m.MPlayer = mPlayer.(MusicPlayer.Model)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	return lipg.JoinHorizontal(
		lipg.Top,
		m.MList.View(),
		m.MPlayer.View(),
	)
}

func Application() {
	mainModel := MainModel{
		MList:   MusicList.New(),
		MPlayer: MusicPlayer.New(),
	}

	bzone.NewGlobal()
	p := tea.NewProgram(
		mainModel,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

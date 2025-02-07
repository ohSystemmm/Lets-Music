package Music

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"io/ioutil"
)

var (
	player *audio.Player
	paused bool
)

func PlaySong(filePath string, audioContext *audio.Context) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	stream, err := wav.Decode(audioContext, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("error decoding file: %w", err)
	}

	player, err = audio.NewPlayer(audioContext, stream)
	if err != nil {
		return fmt.Errorf("error creating player: %w", err)
	}

	player.Play()
	paused = false

	return nil
}

func PauseSong() {
	if paused {
		player.Play()
	} else {
		player.Pause()
	}
	paused = !paused
}

func StopSong() {
	if player != nil {
		player.Close()
	}
}

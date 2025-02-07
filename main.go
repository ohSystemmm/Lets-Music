package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var (
	audioContext *audio.Context
	player       *audio.Player
	paused       bool
)

func main() {
	songDir := "./Backend/songs"

	files, err := ioutil.ReadDir(songDir)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	// Initialize the audio context
	audioContext = audio.NewContext(44100)

	reader := bufio.NewReader(os.Stdin)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := songDir + "/" + file.Name()
		fmt.Printf("Playing: %s\n", filePath)

		err := playSong(filePath)
		if err != nil {
			log.Fatalf("Error playing song: %v", err)
		}

		// Wait for user commands
		for {
			fmt.Println("Enter command (p: pause/resume, s: stop, n: next song):")
			input, _ := reader.ReadString('\n')

			switch input {
			case "p\n":
				pauseSong()
			case "s\n":
				stopSong()
				goto NextSong
			case "n\n":
				stopSong()
				goto NextSong
			default:
				fmt.Println("Invalid command")
			}
		}
	NextSong:
		continue
	}
}

func playSong(filePath string) error {
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

func pauseSong() {
	if paused {
		player.Play()
	} else {
		player.Pause()
	}
	paused = !paused
}

func stopSong() {
	if player != nil {
		player.Close()
	}
}

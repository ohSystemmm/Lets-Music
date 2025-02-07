package main

import (
	"Melodex/Backend/Music"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

func main() {
	songDir := "./Backend/songs"

	files, err := ioutil.ReadDir(songDir)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	audioContext := audio.NewContext(44100)

	reader := bufio.NewReader(os.Stdin)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := songDir + "/" + file.Name()
		fmt.Printf("Playing: %s\n", filePath)

		err := Music.PlaySong(filePath, audioContext)
		if err != nil {
			log.Fatalf("Error playing song: %v", err)
		}

		for {
			fmt.Println("Enter command (p: pause/resume, s: stop, n: next song):")
			input, _ := reader.ReadString('\n')

			switch input {
			case "p\n":
				Music.PauseSong()
			case "s\n":
				Music.StopSong()
				break
			case "n\n":
				Music.StopSong()
				break
			default:
				fmt.Println("Invalid command")
			}

			if input == "s\n" || input == "n\n" {
				break
			}
		}
	}
}

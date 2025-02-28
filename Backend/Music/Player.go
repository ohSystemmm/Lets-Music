package Music

import (
	"bufio"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func PlaySongs(mp3Files []string, context *oto.Context) {
	commandChan := make(chan string)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Error reading input: %v", err)
				continue
			}
			input = strings.TrimSpace(input)
			commandChan <- input
		}
	}()

	for _, file := range mp3Files {
		fmt.Printf("Now playing: %s\n", filepath.Base(file))
		fmt.Println("Press 'n' and hit Enter to skip to the next song.")

		f, err := os.Open(file)
		if err != nil {
			log.Fatalf("Failed to open file %s: %v", file, err)
		}

		decoder, err := mp3.NewDecoder(f)
		if err != nil {
			log.Fatalf("Failed to create MP3 decoder: %v", err)
		}

		player := context.NewPlayer()
		defer player.Close()

		done := make(chan bool)
		stop := make(chan bool)

		go func() {
			buf := make([]byte, 4096)
			for {
				select {
				case <-stop:
					done <- true
					return
				default:
					n, err := decoder.Read(buf)
					if err == io.EOF {
						done <- true
						return
					}
					if err != nil {
						log.Printf("Error reading audio data: %v", err)
						done <- true
						return
					}
					if n > 0 {
						if _, err := player.Write(buf[:n]); err != nil {
							log.Printf("Error playing audio: %v", err)
							done <- true
							return
						}
					}
				}
			}
		}()

		songFinished := false
		for !songFinished {
			select {
			case <-done:
				songFinished = true
			case cmd := <-commandChan:
				switch cmd {
				case "n":
					stop <- true
					<-done
					songFinished = true
				default:
					fmt.Println("Unknown command. Press 'n' to skip.")
				}
			}
		}

		player.Close()
		f.Close()
	}

	fmt.Println("Playback finished.")
}

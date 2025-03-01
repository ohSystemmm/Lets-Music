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
	"sync"
)

type MusicPlayer struct {
	context       *oto.Context
	commandChan   chan string
	mu            sync.Mutex
	paused        bool
	currentPlayer *oto.Player
	currentFile   *os.File
	done          chan bool
	stop          chan bool
}

func NewMusicPlayer(context *oto.Context) *MusicPlayer {
	return &MusicPlayer{
		context:     context,
		commandChan: make(chan string),
	}
}

func (mp *MusicPlayer) PlaySongs(mp3Files []string) {
	go mp.listenForCommands()

	for _, file := range mp3Files {
		mp.playTrack(file)
	}

	fmt.Println("Playback finished.")
}

func (mp *MusicPlayer) playTrack(file string) {
	fmt.Printf("Now playing: %s\n", filepath.Base(file))
	fmt.Println("Press 'n' and hit Enter to skip to the next song. Press 'p' to pause/resume. Press 's' to stop playback.")

	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", file, err)
	}
	defer f.Close()

	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		log.Fatalf("Failed to create MP3 decoder: %v", err)
	}

	player := mp.context.NewPlayer()
	mp.currentPlayer = player
	mp.currentFile = f
	mp.done = make(chan bool)
	mp.stop = make(chan bool)
	defer player.Close()

	go mp.playAudio(decoder)
	mp.waitForCommands()
}

func (mp *MusicPlayer) playAudio(decoder *mp3.Decoder) {
	buf := make([]byte, 4096)
	for {
		select {
		case <-mp.stop:
			mp.done <- true
			return
		default:
			mp.mu.Lock()
			if mp.paused {
				mp.mu.Unlock()
				continue
			}
			mp.mu.Unlock()
			n, err := decoder.Read(buf)
			if err == io.EOF {
				mp.done <- true
				return
			}
			if err != nil {
				log.Printf("Error reading audio data: %v", err)
				mp.done <- true
				return
			}
			if n > 0 {
				if _, err := mp.currentPlayer.Write(buf[:n]); err != nil {
					log.Printf("Error playing audio: %v", err)
					mp.done <- true
					return
				}
			}
		}
	}
}

func (mp *MusicPlayer) listenForCommands() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
			continue
		}
		input = strings.TrimSpace(input)
		mp.commandChan <- input
	}
}

func (mp *MusicPlayer) waitForCommands() {
	songFinished := false
	for !songFinished {
		select {
		case <-mp.done:
			songFinished = true
		case cmd := <-mp.commandChan:
			switch cmd {
			case "n":
				mp.NextTrack()
				songFinished = true
			case "p":
				mp.PauseTrack()
			case "s":
				mp.Stop()
				songFinished = true
			default:
				fmt.Println("Unknown command. Press 'n' to skip, 'p' to pause/resume, 's' to stop.")
			}
		}
	}
}

func (mp *MusicPlayer) PauseTrack() {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	mp.paused = !mp.paused
	if mp.paused {
		fmt.Println("Playback paused.")
	} else {
		fmt.Println("Playback resumed.")
	}
}

func (mp *MusicPlayer) NextTrack() {
	mp.stop <- true
	<-mp.done
	fmt.Println("Skipping to next track.")
}

func (mp *MusicPlayer) Stop() {
	mp.stop <- true
	<-mp.done
	fmt.Println("Playback stopped.")
}

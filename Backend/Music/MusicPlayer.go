package Music

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"os/exec"
	"sync"
	"time"
)

var (
	ctrl  *beep.Ctrl
	mutex sync.Mutex
)

func playMusic() {
	mutex.Lock()
	defer mutex.Unlock()

	cmd := exec.Command("ffmpeg", "-i", "Backend/TestMusic/TestMusic.m4a", "-f", "wav", "pipe:1")
	ffmpegOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Could not create ffmpeg output pipe:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Could not start ffmpeg:", err)
		return
	}

	streamer, format, err := wav.Decode(ffmpegOut)
	if err != nil {
		fmt.Println("Could not decode WAV stream:", err)
		return
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	ctrl = &beep.Ctrl{Streamer: streamer, Paused: false}
	speaker.Play(ctrl)

	fmt.Println("Playing music locally...")
}

func pauseMusic() {
	mutex.Lock()
	defer mutex.Unlock()
	if ctrl != nil {
		ctrl.Paused = !ctrl.Paused
		state := "paused"
		if !ctrl.Paused {
			state = "resumed"
		}
		fmt.Printf("Music %s locally\n", state)
	} else {
		fmt.Println("No music playing")
	}
}

func stopMusic() {
	mutex.Lock()
	defer mutex.Unlock()
	if ctrl != nil {
		ctrl.Streamer = nil
		ctrl = nil
		fmt.Println("Music stopped locally")
	} else {
		fmt.Println("No music playing")
	}
}

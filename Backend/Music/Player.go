package Music

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"os"
	"time"
)

var streamer beep.StreamSeekCloser
var done chan bool
var paused bool

func PlaySong(filePath string, sampleRate beep.SampleRate) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	streamer, format, err := wav.Decode(file)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/9))

	fmt.Println("Playing:", filePath)
	done = make(chan bool)
	paused = false
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	fmt.Println("Finished playing:", filePath)
}

func PauseSong() {
	fmt.Println("Pausing/Resuming...")
	speaker.Lock()
	if paused {
		speaker.Play(streamer)
	} else {
		speaker.Clear()
	}
	paused = !paused
	speaker.Unlock()
}

func StopSong() {
	fmt.Println("Stopping...")
	speaker.Lock()
	if streamer != nil {
		streamer.Close()
	}
	speaker.Unlock()
	if done != nil {
		done <- true
	}
}

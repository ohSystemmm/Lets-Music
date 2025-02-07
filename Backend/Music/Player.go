package Music

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"os"
	"time"
)

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
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	fmt.Println("Finished playing:", filePath)
}

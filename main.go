package main

import (
	"Melodex/Backend/Music"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"io/ioutil"
	"time"
)

func main() {
	songDir := "./Backend/songs"

	files, err := ioutil.ReadDir(songDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	sampleRate := beep.SampleRate(44100)
	speaker.Init(sampleRate, sampleRate.N(time.Second/10))

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := songDir + "/" + file.Name()
		Music.PlaySong(filePath, sampleRate)
	}
}

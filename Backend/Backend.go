package Backend

import (
	"Melodex/Backend/Music"
	"fmt"
	"github.com/hajimehoshi/oto"
	"log"
	"os"
)

func Run() {
	sourceDir := "./Backend/Music/source_songs"
	destDir := "./Backend/Music/mp3_songs"

	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.Mkdir(destDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	mp3Files, err := Music.ConvertToMP3(sourceDir, destDir)
	if err != nil {
		log.Fatalf("Error converting files: %v", err)
	}

	if len(mp3Files) == 0 {
		fmt.Println("No files were converted.")
		return
	}

	context, err := oto.NewContext(44100, 2, 2, 65536)
	if err != nil {
		log.Fatalf("Failed to create audio context: %v", err)
	}
	defer context.Close()

	Music.PlaySongs(mp3Files, context)
}
